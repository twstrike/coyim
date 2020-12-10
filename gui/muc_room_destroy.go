package gui

import (
	"errors"

	"github.com/coyim/coyim/coylog"

	"github.com/coyim/coyim/i18n"
	"github.com/coyim/coyim/xmpp/jid"
	"github.com/coyim/gotk3adapter/gtki"
)

func (v *roomView) onDestroyRoom() {
	d := v.newRoomDestroyView()
	d.show()
}

var (
	errEmptyServiceName   = errors.New("empty service name")
	errEmptyRoomName      = errors.New("empty room name")
	errInvalidRoomName    = errors.New("invalid room name")
	errInvalidServiceName = errors.New("invalid service name")
)

type roomDestroyView struct {
	builder               *builder
	chatServicesComponent *chatServicesComponent
	destroyRoom           func(reason string, alternativeID jid.Bare, password string, done func())

	dialog               gtki.Dialog      `gtk-widget:"destroy-room-dialog"`
	reasonEntry          gtki.TextView    `gtk-widget:"destroy-room-reason-entry"`
	alternativeRoomCheck gtki.CheckButton `gtk-widget:"destroy-room-alternative-check"`
	alternativeRoomBox   gtki.Box         `gtk-widget:"destroy-room-alternative-box"`
	alternativeRoomEntry gtki.Entry       `gtk-widget:"destroy-room-name-entry"`
	passwordEntry        gtki.Entry       `gtk-widget:"destroy-room-password-entry"`
	destroyRoomButton    gtki.Button      `gtk-widget:"destroy-room-button"`
	notificationBox      gtki.Box         `gtk-widget:"notification-area"`

	notification *notifications
}

func (v *roomView) newRoomDestroyView() *roomDestroyView {
	d := &roomDestroyView{}

	d.initBuilder()
	d.initDestroyContext(v)
	d.initChatServices(v)
	d.initDefaults(v)

	return d
}

func (d *roomDestroyView) initBuilder() {
	d.builder = newBuilder("MUCRoomDestroyDialog")
	panicOnDevError(d.builder.bindObjects(d))

	d.builder.ConnectSignals(map[string]interface{}{
		"on_alternative_room_toggled": d.onAlternativeRoomToggled,
		"on_destroy":                  doOnlyOnceAtATime(d.onDestroyRoom),
		"on_cancel":                   d.close,
	})
}

func (d *roomDestroyView) initDestroyContext(v *roomView) {
	d.destroyRoom = func(reason string, alternativeID jid.Bare, password string, done func()) {
		d.close()
		ctx := v.newDestroyContext(reason, alternativeID, password, done)
		ctx.destroyRoom()
	}
}

func (d *roomDestroyView) initChatServices(v *roomView) {
	chatServicesList := d.builder.get("chat-services-list").(gtki.ComboBoxText)
	chatServicesEntry := d.builder.get("chat-services-entry").(gtki.Entry)
	d.chatServicesComponent = v.u.createChatServicesComponent(chatServicesList, chatServicesEntry, nil)
	go d.chatServicesComponent.updateServicesBasedOnAccount(v.account)
}

func (d *roomDestroyView) initDefaults(v *roomView) {
	d.dialog.SetTransientFor(v.window)

	d.notification = v.u.newNotifications(d.notificationBox)
}

// onDestroyRoom MUST be called from the UI thread
func (d *roomDestroyView) onDestroyRoom(done func()) {
	d.notification.clearErrors()

	b, _ := d.reasonEntry.GetBuffer()
	reason := b.GetText(b.GetStartIter(), b.GetEndIter(), false)

	alternativeID, password, err := d.alternativeRoomInformation()
	if err != nil {
		d.notification.error(d.friendlyMessageForAlternativeRoomError(err))
		done()
		return
	}

	d.destroyRoom(reason, alternativeID, password, done)
}

func (d *roomDestroyView) alternativeRoomInformation() (jid.Bare, string, error) {
	if !d.alternativeRoomCheck.GetActive() {
		return nil, "", nil
	}

	alternativeID, err := d.tryParseAlternativeRoomID()
	if err != nil {
		return nil, "", err
	}

	password, _ := d.passwordEntry.GetText()

	return alternativeID, password, nil
}

// onAlternativeRoomToggled MUST be called from the UI thread
func (d *roomDestroyView) onAlternativeRoomToggled() {
	v := d.alternativeRoomCheck.GetActive()
	d.alternativeRoomBox.SetVisible(v)
	d.resetAlternativeRoomFields()
}

func (d *roomDestroyView) resetAlternativeRoomFields() {
	d.alternativeRoomEntry.SetText("")
	d.passwordEntry.SetText("")
	d.chatServicesComponent.resetToDefault()
}

func (d *roomDestroyView) friendlyMessageForAlternativeRoomError(err error) string {
	switch err {
	case errEmptyServiceName:
		return i18n.Local("You must provide a service name")
	case errEmptyRoomName:
		return i18n.Local("You must provide a room name")
	case errInvalidRoomName:
		return i18n.Local("You must provide a valid room name")
	case errInvalidServiceName:
		return i18n.Local("You must provide a valid service name")
	default:
		return i18n.Local("You must provide a valid service and room name")
	}
}

// tryParseAlternativeRoomID MUST be called from the UI thread
//
// This should be "alternative venue" as the protocol says, but
// we prefer to use "alternative room id" in this context
// in order to have a better understanding of what this field means
func (d *roomDestroyView) tryParseAlternativeRoomID() (jid.Bare, error) {
	rn, _ := d.alternativeRoomEntry.GetText()
	s := d.chatServicesComponent.currentServiceValue()

	// We don't really need to continue if the user hasn't entered
	// anything in the room name and the service, because the alternative
	// room is always optional according to the protocol
	if rn == "" && s == "" {
		return nil, nil
	}

	r, err := d.alternativeRoomID()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (d *roomDestroyView) alternativeRoomID() (r jid.Bare, err error) {
	l, err := d.validateRoomName()
	if err != nil {
		return
	}

	s, err := d.validateServiceName()
	if err != nil {
		return
	}

	r = jid.NewBare(l, s)
	return
}

func (d *roomDestroyView) validateRoomName() (l jid.Local, err error) {
	rn, _ := d.alternativeRoomEntry.GetText()

	if rn == "" {
		err = errEmptyRoomName
		return
	}

	l = jid.NewLocal(rn)
	if !l.Valid() {
		err = errInvalidRoomName
	}

	return
}

func (d *roomDestroyView) validateServiceName() (s jid.Domain, err error) {
	if !d.chatServicesComponent.hasServiceValue() {
		err = errEmptyServiceName
		return
	}

	s = jid.NewDomain(d.chatServicesComponent.currentServiceValue())
	if !s.Valid() {
		err = errInvalidServiceName
	}

	return
}

// disableFields MUST be called from the UI thread
func (d *roomDestroyView) disableFields() {
	d.setSensitivityForAllFields(false)
}

// enableFields MUST be called from the UI thread
func (d *roomDestroyView) enableFields() {
	d.setSensitivityForAllFields(true)
}

// setSensitivityForAllFields MUST be called from the UI thread
func (d *roomDestroyView) setSensitivityForAllFields(v bool) {
	d.reasonEntry.SetSensitive(v)
	d.alternativeRoomEntry.SetSensitive(v)
	d.destroyRoomButton.SetSensitive(v)
}

// show MUST be called from the UI thread
func (d *roomDestroyView) show() {
	d.dialog.Show()
}

// close MUST be called from the UI thread
func (d *roomDestroyView) close() {
	d.dialog.Destroy()
}

type roomDestroyContext struct {
	roomID        jid.Bare
	reason        string
	alternativeID jid.Bare
	password      string
	destroy       func(reason string, alternativeID jid.Bare, password string, onSuccess func(), onError func(error), onDone func())
	onDone        func()
	log           coylog.Logger
}

func (v *roomView) newDestroyContext(reason string, alternativeID jid.Bare, password string, done func()) *roomDestroyContext {
	return &roomDestroyContext{
		roomID:        v.roomID(),
		reason:        reason,
		alternativeID: alternativeID,
		password:      password,
		destroy:       v.tryDestroyRoom,
		onDone:        done,
		log:           v.log,
	}
}

func (dc *roomDestroyContext) destroyRoom() {
	dc.destroy(dc.reason, dc.alternativeID, dc.password, dc.onDestroySuccess, dc.onDestroyFails, dc.onDone)
}

func (dc *roomDestroyContext) onDestroySuccess() {
	dc.log.Info("The room has been destroyed")
}

func (dc *roomDestroyContext) onDestroyFails(err error) {
	doInUIThread(func() {
		rd := newDestroyError(dc.roomID, err, dc.destroyRoom)
		rd.show()
	})
}

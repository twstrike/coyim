package gui

import (
	"sync"

	"github.com/coyim/coyim/session/muc"
	"github.com/coyim/coyim/xmpp/jid"
	"github.com/coyim/gotk3adapter/gtki"
)

// TODO: message bar (for errors etc)
// TODO: it should be possible to put in your own custom chat service as well
// TODO: add refresh button

const (
	mucListRoomsIndexJid         = 0
	mucListRoomsIndexName        = 1
	mucListRoomsIndexService     = 2
	mucListRoomsIndexDescription = 3
	mucListRoomsIndexOccupants   = 4
)

type roomListingUpdateData struct {
	iter       gtki.TreeIter
	view       *mucPublicRoomsView
	generation int
}

func (u *gtkUI) updatedRoomListing(rl *muc.RoomListing, data interface{}) {
	d := data.(*roomListingUpdateData)

	// If we get an old update, we don't want to do anything at all
	if d.view.generation == d.generation {
		doInUIThread(func() {
			_ = d.view.roomsModel.SetValue(d.iter, mucListRoomsIndexDescription, rl.Description)
			_ = d.view.roomsModel.SetValue(d.iter, mucListRoomsIndexOccupants, rl.Occupants)
		})
	}
}

type mucPublicRoomsView struct {
	builder *builder

	generation    int
	updateLock    sync.RWMutex
	serviceGroups map[string]gtki.TreeIter
	cancel        chan bool

	dialog       gtki.Dialog         `gtk-widget:"MUCPublicRooms"`
	model        gtki.ListStore      `gtk-widget:"accounts-model"`
	accountInput gtki.ComboBox       `gtk-widget:"accounts"`
	roomsModel   gtki.TreeStore      `gtk-widget:"rooms-model"`
	roomsTree    gtki.TreeView       `gtk-widget:"rooms-tree"`
	rooms        gtki.ScrolledWindow `gtk-widget:"rooms"`
	spinner      gtki.Spinner        `gtk-widget:"spinner"`

	accountsList    []*account
	accounts        map[string]*account
	currentlyActive int
}

func (prv *mucPublicRoomsView) init() {
	prv.builder = newBuilder("MUCPublicRoomsDialog")
	panicOnDevError(prv.builder.bindObjects(prv))
	prv.serviceGroups = make(map[string]gtki.TreeIter)
}

// initOrReplaceAccounts should be called from the UI thread
func (prv *mucPublicRoomsView) initOrReplaceAccounts(accounts []*account) {
	// TODO: if there are no active accounts, maybe we should just hide the spinner and the view

	prv.currentlyActive = 0
	if prv.accounts != nil {
		act := prv.accountInput.GetActive()
		currentlyActiveAccount := prv.accountsList[act]
		for ix, a := range accounts {
			if currentlyActiveAccount == a {
				prv.currentlyActive = ix
			}
		}
		prv.model.Clear()
	}

	prv.accountsList = accounts
	prv.accounts = make(map[string]*account)
	for _, acc := range accounts {
		iter := prv.model.Append()
		_ = prv.model.SetValue(iter, 0, acc.session.GetConfig().Account)
		_ = prv.model.SetValue(iter, 1, acc.session.GetConfig().ID())
		prv.accounts[acc.session.GetConfig().ID()] = acc
	}

	if len(accounts) > 0 {
		prv.accountInput.SetActive(prv.currentlyActive)
	}
}

// mucUpdatePublicRoomsOn should NOT be called from the UI thread
func (u *gtkUI) mucUpdatePublicRoomsOn(view *mucPublicRoomsView, a *account) {
	if view.cancel != nil {
		view.cancel <- true
	}

	view.updateLock.Lock()

	view.cancel = make(chan bool, 1)

	doInUIThread(func() {
		view.rooms.SetVisible(false)
		view.spinner.Start()
		view.spinner.SetVisible(true)
		view.roomsModel.Clear()
	})
	view.generation++
	view.serviceGroups = make(map[string]gtki.TreeIter)

	hasSomething := false

	// We save the generation value here, in case it gets modified inside the view later
	gen := view.generation

	res, resServices, ec := a.session.GetRooms(jid.Parse(a.session.GetConfig().Account).Host())
	go func() {
		defer func() {
			if !hasSomething {
				doInUIThread(func() {
					view.spinner.Stop()
					view.spinner.SetVisible(false)
					view.rooms.SetVisible(true)
				})
			}

			view.updateLock.Unlock()
		}()
		for {
			select {
			case sl, ok := <-resServices:
				if !ok {
					return
				}
				if !hasSomething {
					hasSomething = true
					doInUIThread(func() {
						view.spinner.Stop()
						view.spinner.SetVisible(false)
						view.rooms.SetVisible(true)
					})
				}

				serv, ok := view.serviceGroups[sl.Jid.String()]
				if !ok {
					doInUIThread(func() {
						serv = view.roomsModel.Append(nil)
						view.serviceGroups[sl.Jid.String()] = serv
						_ = view.roomsModel.SetValue(serv, mucListRoomsIndexJid, sl.Jid.String())
						_ = view.roomsModel.SetValue(serv, mucListRoomsIndexName, sl.Name)
					})
				}
			case rl, ok := <-res:
				if !ok || rl == nil {
					return
				}

				if !hasSomething {
					hasSomething = true
					doInUIThread(func() {
						view.spinner.Stop()
						view.spinner.SetVisible(false)
						view.rooms.SetVisible(true)
					})
				}

				serv, ok := view.serviceGroups[rl.Service.String()]
				doInUIThread(func() {
					if !ok {
						serv = view.roomsModel.Append(nil)
						view.serviceGroups[rl.Service.String()] = serv
						_ = view.roomsModel.SetValue(serv, mucListRoomsIndexJid, rl.Service.String())
						_ = view.roomsModel.SetValue(serv, mucListRoomsIndexName, rl.ServiceName)
					}

					iter := view.roomsModel.Append(serv)
					_ = view.roomsModel.SetValue(iter, mucListRoomsIndexJid, string(rl.Jid.Local()))
					_ = view.roomsModel.SetValue(iter, mucListRoomsIndexName, rl.Name)
					_ = view.roomsModel.SetValue(iter, mucListRoomsIndexService, rl.Service.String())
					rl.OnUpdate(u.updatedRoomListing, &roomListingUpdateData{iter, view, gen})

					view.roomsTree.ExpandAll()
				})

			case e, ok := <-ec:
				if !ok {
					return
				}
				if e != nil {
					u.log.WithError(e).Debug("something went wrong trying to get chat rooms")
				}
				return
			case _, _ = <-view.cancel:
				return
			}
		}
	}()
}

// mucShowPublicRooms should be called from the UI thread
func (u *gtkUI) mucShowPublicRooms() {
	view := &mucPublicRoomsView{}
	view.init()

	view.initOrReplaceAccounts(u.getAllConnectedAccounts())

	accountsObserverToken := u.onChangeOfConnectedAccounts(func() {
		doInUIThread(func() {
			view.initOrReplaceAccounts(u.getAllConnectedAccounts())
		})
	})

	view.builder.ConnectSignals(map[string]interface{}{
		"on_cancel_signal": view.dialog.Destroy,
		"on_close_window_signal": func() {
			u.removeConnectedAccountsObserver(accountsObserverToken)
		},
		"on_join_signal": func() {},
		"on_accounts_changed": func() {
			act := view.accountInput.GetActive()
			if act >= 0 && act < len(view.accountsList) && act != view.currentlyActive {
				view.currentlyActive = act
				go u.mucUpdatePublicRoomsOn(view, view.accountsList[act])
			}
		},
	})

	view.dialog.SetTransientFor(u.window)
	view.dialog.Show()
	view.currentlyActive = -1

	if len(view.accountsList) > 0 {
		view.currentlyActive = 0
		go u.mucUpdatePublicRoomsOn(view, view.accountsList[view.currentlyActive])
	}
}
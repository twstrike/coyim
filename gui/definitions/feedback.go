package definitions

func init() {
	add(`Feedback`, &defFeedback{})
}

type defFeedback struct{}

func (*defFeedback) String() string {
	return `
<interface>
  <object class="GtkMessageDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="title" translatable="yes">We would like to receive your feedback</property>
    <property name="modal">true</property>
    <property name="secondary-text" translatable="yes">Please, tell us how is going for you to use CoyIM.&#xA;This is the only way we can create a better tool to keep your conversations private.</property>
    <property name="text" translatable="yes">https://coy.im/feedback</property>
    <property name="buttons">GTK_BUTTONS_CLOSE</property>
  </object>
</interface>

`
}

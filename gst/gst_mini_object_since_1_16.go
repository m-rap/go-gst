//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"

// AddParent adds the given object as a parent of this object.
// See https://gstreamer.freedesktop.org/documentation/gstreamer/gstminiobject.html?gi-language=c#gst_mini_object_add_parent.
func (m *MiniObject) AddParent(parent *MiniObject) {
	C.gst_mini_object_add_parent(m.Instance(), parent.Instance())
}

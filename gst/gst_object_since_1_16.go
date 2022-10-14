//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "unsafe"

// Clear will will clear all references to this object. If the reference is already null
// the the function does nothing. Otherwise the reference count is decreased and the pointer
// set to null.
func (o *Object) Clear() {
	if ptr := o.Unsafe(); ptr != nil {
		C.gst_clear_object((**C.GstObject)(unsafe.Pointer(&ptr)))
	}
}

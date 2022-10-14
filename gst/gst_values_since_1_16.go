//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import (
	"unsafe"

	"github.com/tinyzimmer/go-glib/glib"
)

// ValueArray converts the given slice of Go types into a ValueArrayValue.
// This function can return nil on any conversion or memory allocation errors.
func ValueArray(ss []interface{}) *ValueArrayValue {
	v, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_array_init(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(len(ss)),
	)
	for _, s := range ss {
		val, err := glib.GValue(s)
		if err != nil {
			return nil
		}
		C.gst_value_array_append_value(
			(*C.GValue)(unsafe.Pointer(v.GValue)),
			(*C.GValue)(unsafe.Pointer(val.GValue)),
		)
	}
	out := ValueArrayValue(*v)
	return &out
}

// ValueList converts the given slice of Go types into a ValueListValue.
// This function can return nil on any conversion or memory allocation errors.
func ValueList(ss []interface{}) *ValueListValue {
	v, err := glib.ValueAlloc()
	if err != nil {
		return nil
	}
	C.gst_value_list_init(
		(*C.GValue)(unsafe.Pointer(v.GValue)),
		C.guint(len(ss)),
	)
	for _, s := range ss {
		val, err := glib.GValue(s)
		if err != nil {
			return nil
		}
		C.gst_value_list_append_value(
			(*C.GValue)(unsafe.Pointer(v.GValue)),
			(*C.GValue)(unsafe.Pointer(val.GValue)),
		)
	}
	out := ValueListValue(*v)
	return &out
}

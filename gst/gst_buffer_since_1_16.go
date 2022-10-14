//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "unsafe"

// NewBufferFromBytes returns a new buffer from the given byte slice.
func NewBufferFromBytes(b []byte) *Buffer {
	gbytes := C.g_bytes_new((C.gconstpointer)(unsafe.Pointer(&b[0])), C.gsize(len(b)))
	defer C.g_bytes_unref(gbytes)
	return FromGstBufferUnsafeFull(unsafe.Pointer(C.gst_buffer_new_wrapped_bytes(gbytes)))
}

// NewBufferFromReader returns a new buffer from the given io.Reader.
func NewBufferFromReader(rdr io.Reader) (*Buffer, error) {
	out, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}
	return NewBufferFromBytes(out), nil
}

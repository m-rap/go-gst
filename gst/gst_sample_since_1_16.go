//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"

// SetBuffer sets the buffer inside this sample. The sample must be writable.
func (s *Sample) SetBuffer(buf *Buffer) {
	C.gst_sample_set_buffer(s.Instance(), buf.Instance())
}

// SetCaps sets the caps on this sample. The sample must be writable.
func (s *Sample) SetCaps(caps *Caps) { C.gst_sample_set_caps(s.Instance(), caps.Instance()) }

// SetInfo sets the info on this sample. The sample must be writable.
func (s *Sample) SetInfo(st *Structure) bool {
	return gobool(C.gst_sample_set_info(s.Instance(), st.Instance()))
}

// SetSegment sets the segment on this sample. The sample must be writable.
func (s *Sample) SetSegment(segment *Segment) {
	C.gst_sample_set_segment(s.Instance(), segment.Instance())
}

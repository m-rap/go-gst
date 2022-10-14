//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "unsafe"

// NewBitrateQuery constructs a new query object for querying the bitrate.
func NewBitrateQuery() *Query {
	return FromGstQueryUnsafeFull(unsafe.Pointer(C.gst_query_new_bitrate()))
}

// ParseBitrate gets the results of a bitrate query. See also SetBitrate.
func (q *Query) ParseBitrate() uint {
	var out C.guint
	C.gst_query_parse_bitrate(q.Instance(), &out)
	return uint(out)
}

// SetBitrate sets the results of a bitrate query. The nominal bitrate is the average bitrate expected over the length of the stream as advertised
// in file headers (or similar).
func (q *Query) SetBitrate(nominal uint) {
	C.gst_query_set_bitrate(q.Instance(), C.guint(nominal))
}

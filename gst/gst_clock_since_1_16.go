//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "unsafe"

// GetClock returns the clock for this ClockID.
func (c *ClockID) GetClock() *Clock {
	clk := C.gst_clock_id_get_clock(c.Instance())
	return FromGstClockUnsafeFull(unsafe.Pointer(clk))
}

// UsesClock returns whether id uses clock as the underlying clock. clock can be nil, in which case the return value indicates whether the
// underlying clock has been freed. If this is the case, the id is no longer usable and should be freed.
func (c *ClockID) UsesClock(clock *Clock) bool {
	return gobool(C.gst_clock_id_uses_clock(c.Instance(), clock.Instance()))
}

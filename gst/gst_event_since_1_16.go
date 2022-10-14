//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "time"

// ParseSeekTrickModeInterval retrieves the trickmode interval that may have been set on a seek event with
// SetSeekTrickModeInterval.
func (e *Event) ParseSeekTrickModeInterval() (interval time.Duration) {
	var out C.GstClockTime
	C.gst_event_parse_seek_trickmode_interval(e.Instance(), &out)
	return time.Duration(out)
}

// SetSeekTrickModeInterval sets a trickmode interval on a (writable) seek event. Elements that support TRICKMODE_KEY_UNITS
// seeks SHOULD use this as the minimal interval between each frame they may output.
func (e *Event) SetSeekTrickModeInterval(interval time.Duration) {
	C.gst_event_set_seek_trickmode_interval(e.Instance(), C.GstClockTime(interval.Nanoseconds()))
}

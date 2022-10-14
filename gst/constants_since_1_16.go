//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"

const (
	MessageDeviceChanged MessageType = C.GST_MESSAGE_DEVICE_CHANGED
	QueryBitrate         QueryType   = C.GST_QUERY_BITRATE // (51202) – the bitrate query (since 1.16)
)

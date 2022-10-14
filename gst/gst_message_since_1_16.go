//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"

import (
	"fmt"
	"unsafe"
)

// ParseDeviceChanged Parses a device-changed message. The device-changed message is
// produced by GstDeviceProvider or a GstDeviceMonitor. It announces that a device's properties
// have changed.
// The first device returned is the updated Device, and the second changedDevice represents
// the old state of the device.
func (m *Message) ParseDeviceChanged() (newDevice, oldDevice *Device) {
	var gstNewDevice, gstOldDevice *C.GstDevice
	C.gst_message_parse_device_changed((*C.GstMessage)(m.Instance()), &gstNewDevice, &gstOldDevice)
	return FromGstDeviceUnsafeFull(unsafe.Pointer(gstNewDevice)),
		FromGstDeviceUnsafeFull(unsafe.Pointer(gstOldDevice))
}

type Message_since_1_16 struct {
	Message
}

func (m *Message_since_1_16) String() string {
	msg := m.Message.String()
	switch m.Type() {
	case MessageDeviceChanged:
		if device, _ := m.ParseDeviceChanged(); device != nil {
			msg += fmt.Sprintf("Device %s had its properties updated", device.GetDisplayName())
		}
	}
	return msg
}

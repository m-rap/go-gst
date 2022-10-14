//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

import "C"
import "unsafe"

// NewDeviceChangedMessage creates a new device-changed message. The device-changed message is produced by a DeviceProvider or a DeviceMonitor.
// They announce that a device properties has changed and device represent the new modified version of changed_device.
func NewDeviceChangedMessage(src interface{}, device, changedDevice *Device) *Message {
	srcObj := getMessageSourceObj(src)
	if srcObj == nil {
		return nil
	}
	return FromGstMessageUnsafeFull(unsafe.Pointer(C.gst_message_new_device_changed(srcObj, device.Instance(), changedDevice.Instance())))
}

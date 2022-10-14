//go:build !gstreamer_1_14
// +build !gstreamer_1_14

package gst

// SetFeaturesSimple sets the CapsFeatures for all the structures of these caps.
func (c *Caps) SetFeaturesSimple(features *CapsFeatures) {
	if features == nil {
		C.gst_caps_set_features_simple(c.Instance(), nil)
		return
	}
	C.gst_caps_set_features_simple(c.Instance(), features.Instance())
}

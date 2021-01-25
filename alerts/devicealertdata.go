package alerts

// DeviceAlertData represents the device data in alert emails
type DeviceAlertData struct {
	DeviceID   string
	Properties []PropertyAlertData
}

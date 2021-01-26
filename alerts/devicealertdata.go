package alerts

// DeviceAlertData represents the device data in alert emails
type DeviceAlertData struct {
	ID         string
	Name       string
	Properties []PropertyAlertData
}

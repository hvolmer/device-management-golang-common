package alerts

// DeviceAlertData represents the device data in alert emails
type DeviceAlertData struct {
	ID         string `bson:"_id"`
	Name       string
	Properties []PropertyAlertData
}

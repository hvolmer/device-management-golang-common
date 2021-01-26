package alerts

// SystemAlertData represents the system and company data in queued
// alert messages
type SystemAlertData struct {
	CompanyUUID string
	CompanyName string
	SystemUUID  string
	SystemName  string
	Devices     []DeviceAlertData
}

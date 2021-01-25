package alerts

// SystemAlertData represents the system and company data in queued
// alert messages
type SystemAlertData struct {
	CompanyUUID string
	SystemUUID  string
	Devices     []DeviceAlertData
}

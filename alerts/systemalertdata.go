package alerts

// SystemAlertData represents the system and company data in queued
// alert messages
type SystemAlertData struct {
	CompanyUUID    string
	CompanyName    string
	SystemLocation PortalSystemLocation
	SystemUUID     string
	SystemName     string
	Devices        []DeviceAlertData
}

// PortalSystemLocation ...
type PortalSystemLocation struct {
	Location string `json:"location" bson:"location"`
	Building string `json:"building" bson:"building"`
	Floor    string `json:"floor" bson:"floor"`
	Room     string `json:"room" bson:"room"`
}

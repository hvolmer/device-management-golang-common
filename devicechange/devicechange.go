package devicechange

// DeviceChange represents the database type for change logging.
type DeviceChange struct {
	Comment     string      `json:"comment,omitempty" bson:"comment,omitempty"`
	CompanyUUID string      `json:"companyUuid" bson:"companyUuid"`
	Compliance  string      `json:"compliance,omitempty" bson:"compliance,omitempty"`
	DeviceID    string      `json:"deviceId" bson:"deviceId"`
	Path        string      `json:"path" bson:"path"`
	Timestamp   string      `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Type        string      `json:"type" bson:"type"`
	User        string      `json:"user,omitempty" bson:"user,omitempty"`
	Value       interface{} `json:"value" bson:"value"`
}

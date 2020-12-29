package propchange

import "time"

// PropertyChangeMessage ...
type PropertyChangeMessage struct {
	Changes     []PropertyChange `json:"changes"`
	CompanyUUID string           `json:"companyUuid"`
	DeviceID    string           `json:"deviceId"`
	Time        time.Time        `json:"time"`
}

// PropertyChange ...
type PropertyChange struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

package devicechange

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hvolmer/device-management-golang-common/device"
)

// PropertyChangeMessage ...
type PropertyChangeMessage struct {
	Changes         []PropertyChange `json:"changes"`
	CompanyUUID     string           `json:"companyUuid"`
	ComplianceItems []device.ComplianceItem
	DeviceID        primitive.ObjectID `json:"deviceId"`
	Time            time.Time          `json:"time"`
}

// PropertyChange ...
type PropertyChange struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

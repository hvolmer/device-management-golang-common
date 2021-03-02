package alerts

import (
	"time"

	"github.com/hvolmer/device-management-golang-common/devicechange"
)

// PropertyAlertData ...
type PropertyAlertData struct {
	Path          string
	NameWithGroup string
	Label         string
	// Last State one of fault, recovery, schema-fault, schema-recovery
	LastState      string
	CurrentChange  devicechange.MinimizedDeviceChange
	UserChange     devicechange.MinimizedDeviceChange
	UserChangeName string
	PreviousChange devicechange.MinimizedDeviceChange
	Priority       int
	Timeout        time.Time
	AlertedTime    time.Time
}

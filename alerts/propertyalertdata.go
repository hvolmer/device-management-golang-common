package alerts

import "github.com/hvolmer/device-management-golang-common/devicechange"

// PropertyAlertData ...
type PropertyAlertData struct {
	Path                   string
	CurrentChange          devicechange.MinimizedDeviceChange
	UserChange             devicechange.MinimizedDeviceChange
	PreviousOppositeChange devicechange.MinimizedDeviceChange
}

package alerts

import "github.com/hvolmer/device-management-golang-common/devicechange"

// PropertyAlertData ...
type PropertyAlertData struct {
	Path                   string
	NameWithGroup          string
	Label                  string
	LastState              string
	CurrentChange          devicechange.MinimizedDeviceChange
	UserChange             devicechange.MinimizedDeviceChange
	UserChangeName         string
	PreviousOppositeChange devicechange.MinimizedDeviceChange
}

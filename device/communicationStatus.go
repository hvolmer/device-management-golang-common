package device

import "time"

// CommunicationStatus is used to hold all values necessary to compute
// the various online states of a device.
type CommunicationStatus struct {
	CanPing      bool
	CanPoll      bool
	OfflineTime  time.Time
	PingDisabled bool
	PollDisabled bool
	Time         time.Time
}

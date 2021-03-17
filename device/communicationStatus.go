package device

import "time"

// CommunicationStatus is used to hold all values necessary to compute
// the various online states of a device.
type CommunicationStatus struct {
	// CanPing is true when the device can be pinged
	CanPing bool
	// CanPoll is true when the device can be polled
	CanPoll bool
	// IsFirstOnline is true only when the status contains the first
	// transition to online
	IsFirstOnline bool
	// PingDisabled is true when the device, template or type settings disable
	// ping.
	PingDisabled bool
	// PollDisabled is true when the device, template or type settings disable
	// polling.
	PollDisabled bool
	// Time is the time of this message
	Time time.Time
}

// IsOnline checks the various states and disables to calculate whether
// this is considered online.
func (cs CommunicationStatus) IsOnline() bool {
	if !cs.PollDisabled {
		return cs.CanPoll
	}
	if !cs.PingDisabled {
		return cs.CanPing
	}
	return false
}

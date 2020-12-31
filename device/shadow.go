package device

// Shadow ...
type Shadow struct {
	State ShadowState `json:"state" bson:"state"`
}

// ShadowState ...
type ShadowState struct {
	Reported StateReported `json:"reported" bson:"reported"`
}

// StateReported ...
type StateReported struct {
	Error                    string          `json:"error" bson:"error"`
	Online                   bool            `json:"online" bson:"online"`
	OfflineDetectedTimestamp string          `json:"offlineDetectedTimestamp" bson:"offlineDetectedTimestamp"`
	Network                  ReportedNetwork `json:"network" bson:"network"`
}

// ReportedNetwork ...
type ReportedNetwork struct {
	Ping ReportedNetworkPing `json:"reportedNetworkPing" bson:"reportedNetworkPing"`
}

// ReportedNetworkPing ...
type ReportedNetworkPing struct {
	Pingable bool `json:"pingable" bson:"pingable"`
}

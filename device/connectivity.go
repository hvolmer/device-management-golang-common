package device

// Connectivity ...
type Connectivity struct {
	OfflineDetectedTimestamp string `bson:"offlineDetectedTimestamp,omitempty"`
	Online                   bool   `bson:"online"`
	Pingable                 bool   `bson:"pingable"`
	Error                    string `bson:"error,omitempty"`
}

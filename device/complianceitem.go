package device

// ComplianceItem ...
type ComplianceItem struct {
	AlertCount int32            `json:"alertCount" bson:"alertCount"`
	Error      string           `json:"error" bson:"error"`
	HasAlerted bool             `json:"hasAlerted" bson:"hasAlerted"`
	Label      string           `bson:"label"`
	LastValue  interface{}      `bson:"lastValue"`
	Path       string           `json:"path" bson:"path"`
	Schema     ComplianceSchema `json:"schema" bson:"schema"`
}

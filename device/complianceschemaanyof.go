package device

// ComplianceSchemaAnyOf ...
type ComplianceSchemaAnyOf struct {
	Type    string      `json:"type" bson:"type"`
	Const   interface{} `json:"const" bson:"const"`
	Pattern string      `json:"pattern" bson:"pattern"`
}

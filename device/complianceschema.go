package device

// ComplianceSchema ...
type ComplianceSchema struct {
	AnyOf []ComplianceSchemaAnyOf `json:"anyOf" bson:"anyOf"`
	Lock  bool                    `json:"lock" bson:"lock"`
}

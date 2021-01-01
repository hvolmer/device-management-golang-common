package device

import (
	"log"

	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hvolmer/device-management-golang-common/devicechange"
)

// Device type represents the basics of a Device Management device, where the query includes a
// type, not just an ID
type Device struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyUUID     string             `json:"companyUuid" bson:"companyUuid"`
	ComplianceItems []ComplianceItem   `json:"complianceItems" bson:"complianceItems"`
	Connectivity    Connectivity       `bson:"connectivity" bson:"connectivity"`
	Configuration   Configuration      `json:"configuration" bson:"configuration"`
	Shadow          Shadow             `bson:"shadow,omitempty" bson:"shadow"`
	Type            primitive.ObjectID `bson:"type,omitempty" bson:"type"`
}

// GetComplianceForChange ...
func (dev Device) GetComplianceForChange(ch devicechange.DeviceChange) string {
	for _, ci := range dev.ComplianceItems {
		if ci.Path == ch.Path {
			schemaLoader := gojsonschema.NewGoLoader(ci.Schema)
			documentLoader := gojsonschema.NewGoLoader(ch.Value)
			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				log.Println("Error validating compliance. [value]", ch.Value, "[schema]", ci.Schema)
			}
			log.Println("COMPLIANCE RESULT:", result)
		}
	}
	return ""
}

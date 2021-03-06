package device

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

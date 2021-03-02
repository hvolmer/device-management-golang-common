package device

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Device type represents the basics of a Device Management device, where the query includes a
// type, not just an ID
type Device struct {
	ID                  primitive.ObjectID  `json:"_id" bson:"_id"`
	ChangeItems         []ChangeItem        `json:"changeItems" bson:"changeItems"`
	CompanyUUID         string              `json:"companyUuid" bson:"companyUuid"`
	CommunicationStatus CommunicationStatus `json:"communicationStatus" bson:"communicationStatus"`
	ComplianceItems     []ComplianceItem    `json:"complianceItems" bson:"complianceItems"`
	Connectivity        Connectivity        `bson:"connectivity" bson:"connectivity"`
	Configuration       Configuration       `json:"configuration" bson:"configuration"`
	Shadow              Shadow              `bson:"shadow,omitempty" bson:"shadow"`
	Type                primitive.ObjectID  `bson:"type,omitempty" bson:"type"`
}

// WithType represents the basics of a Device Management device, where the query includes a
// type, not just an ID
type WithType struct {
	Device
	// ID              primitive.ObjectID `json:"_id" bson:"_id"`
	// CompanyUUID     string             `json:"companyUuid" bson:"companyUuid"`
	// ComplianceItems []ComplianceItem   `json:"complianceItems" bson:"complianceItems"`
	// Connectivity    Connectivity       `bson:"connectivity" bson:"connectivity"`
	// Configuration   Configuration      `json:"configuration" bson:"configuration"`
	// Shadow          Shadow             `bson:"shadow,omitempty" bson:"shadow"`
	Type Type `bson:"type,omitempty" bson:"type"`
}

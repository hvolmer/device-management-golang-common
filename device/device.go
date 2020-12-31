package device

import "go.mongodb.org/mongo-driver/bson/primitive"

// Device type represents the basics of a Device Management device
type Device struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ComplianceItems []ComplianceItem   `json:"complianceItems" bson:"complianceItems"`
	Connectivity    Connectivity       `bson:"connectivity"`
	Configuration   Configuration      `json:"configuration" bson:"configuration"`
	Shadow          Shadow             `bson:"shadow,omitempty"`
	Type            Type               `bson:"type,omitempty"`
}

package devicechange

import "go.mongodb.org/mongo-driver/bson/primitive"

// MinimizedDeviceChange represents the database type for change logging.
type MinimizedDeviceChange struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Comment    string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Compliance string             `json:"compliance,omitempty" bson:"compliance,omitempty"`
	Type       string             `json:"type" bson:"type"`
	User       string             `json:"user,omitempty" bson:"user,omitempty"`
	Value      interface{}        `json:"value" bson:"value"`
}

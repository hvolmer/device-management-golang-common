package alerts

import "go.mongodb.org/mongo-driver/bson/primitive"

// DeviceAlertData represents the device data in alert emails
type DeviceAlertData struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string
	Properties []PropertyAlertData
}

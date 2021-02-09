package device

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChangeItem contains the state of change operations on a device
type ChangeItem struct {
	ID                      primitive.ObjectID `bson:"_id"`
	Comment                 string             `bson:"comment"`
	ExpiresDate             time.Time          `bson:"expiresDate"`
	PropertyDefinition      primitive.ObjectID `bson:"propertyDefinition"`
	PropertyDefinitionGroup primitive.ObjectID `bson:"propertyDefinitionGroup"`
	Status                  string             `bson:"status"`
	Value                   interface{}        `bson:"value"`
}

// ChangeItemPopulated contains a properly-populated change item
type ChangeItemPopulated struct {
	ID                      primitive.ObjectID      `bson:"_id"`
	Comment                 string                  `bson:"comment"`
	ExpiresDate             time.Time               `bson:"expiresDate"`
	PropertyDefinition      PropertyDefinition      `bson:"propertyDefinition"`
	PropertyDefinitionGroup PropertyDefinitionGroup `bson:"propertyDefinitionGroup"`
	Status                  string                  `bson:"status"`
	Value                   interface{}             `bson:"value"`
}

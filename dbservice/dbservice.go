package dbservice

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBClient ...
var dBClient mongo.Client

// Database is the devices database
var Database *mongo.Database

// AlertMessagesColl The created alerts
var AlertMessagesColl *mongo.Collection

// AlertTemplatesColl ...
var AlertTemplatesColl *mongo.Collection

// DevicesColl The devices collection
var DevicesColl *mongo.Collection

// DeviceChangesColl ...
var DeviceChangesColl *mongo.Collection

// PropertyDefinitionsColl The property-definitions collection
var PropertyDefinitionsColl *mongo.Collection

// DeviceTypesColl The device types collection
var DeviceTypesColl *mongo.Collection

// Start Starts all this up.
func Start() {
	dbURI := os.Getenv("DEVMGMT_DB_URI")
	fmt.Println("GO DB!", dbURI)
	dBClient, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		// fmt.Println("CANCEL???")
	}
	err = dBClient.Connect(ctx)
	if err != nil {
		log.Fatal("ERROR!", err) // ################################################## REALLY A BIG BOOM ERROR

	}
	fmt.Println("DB CONNECTED")
	Database = dBClient.Database("devices")
	AlertMessagesColl = Database.Collection("alertmessages")
	AlertTemplatesColl = Database.Collection("alerttemplates")
	DevicesColl = Database.Collection("devices")
	DeviceChangesColl = Database.Collection("devicechanges")
	DeviceTypesColl = Database.Collection("devicetypes")
}

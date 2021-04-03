package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/noisersup/dashboard-backend-pomodoro/models"
	u "github.com/noisersup/dashboard-backend-pomodoro/utils"
)

type Database struct{
	client *mongo.Client
	coll 	*mongo.Collection
}

func ConnectToDatabase(uri string, dbName string, collName string) (*Database,error){
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err !=nil{ return nil,u.Err("NewClient",err) }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = client.Connect(ctx)
	if err !=nil{ return nil,u.Err("Connect",err) }

	err = client.Ping(ctx, readpref.Primary())
	if err !=nil{ return nil,u.Err("Ping",err) }
 
	coll := client.Database(dbName).Collection(collName)

	return &Database{client,coll},nil
}

func (db *Database) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return db.client.Disconnect(ctx)
}

func (db *Database) GetTimestamp() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var timestamp models.Timestamp

	err := db.coll.FindOne(ctx,  bson.M{"_id": "0"}).Decode(&timestamp)
	
	return timestamp.Timestamp, err
}

func (db *Database) SetTimestamp(timestamp int) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Update().SetUpsert(true)
	_, err := db.coll.UpdateOne(
		ctx,
		bson.M{"_id": "0"},
		bson.D{
			{"$set", bson.D{{"timestamp",timestamp},}},
		},options)

	return err
}
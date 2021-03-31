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

	cur, err := db.coll.Find(ctx,  bson.D{{}}, options.Find())
	if err != nil {return 0,err}
	
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var resp models.TimestampDb

	for cur.Next(ctx){
		err = cur.Decode(&resp)
		if err != nil {return 0,err}
	}
	if err := cur.Err(); err != nil{
		return 0,err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur.Close(ctx)

	return resp.Timestamp, nil
}
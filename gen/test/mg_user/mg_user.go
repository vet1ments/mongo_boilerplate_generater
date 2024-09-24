package mgusertest

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
	"boilerplate/gen/test/mg_embed"
)

var Collection = "User"

type connection struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

type Doc struct {
	ID bson.ObjectID `bson:"_id" json:"id"`	
	Sex string `bson:"sex" json:"sex,omitempty"`
	PhoneNumber mgembedtest.Login `bson:"phone_number" json:"phone_number,omitempty"`
	test mgembedtest.[]string `bson:"test" json:"test,omitempty"`
}

var conn = &connection{}

func IsInit() error {
	if conn.Client == nil || conn.DB == nil || conn.Collection == nil {
		return errors.New("Not initialized")
	}
	return nil
}

func Init(client *mongo.Client, db *mongo.Database) {
	conn.Client = client
	conn.DB = db
	conn.Collection = db.Collection(Collection)
}

func GetUser(id bson.ObjectID) (*Doc, error) {
	if e := IsInit(); e != nil {
		return nil, e
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	model := &Doc{}
	err := conn.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetAllUser() ([]*Doc, error) {
	if e := IsInit(); e != nil {
		return nil, e
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := conn.Collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	models := make([]*Doc, 0)
	for cur.Next(ctx) {
		model := &Doc{}
		if err = cur.Decode(&model); err != nil {
			continue
		} else {
			models = append(models, model)
		}
	}
	return models, nil
}

func Create(model *Doc) (*Doc, error) {
	if e := IsInit(); e != nil {
		return nil, e
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	model.ID = bson.NewObjectID()
	_, err := conn.Collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetCollection() (*mongo.Collection, error) {
	if e := IsInit(); e != nil {
		return nil, e
	}
	return conn.Collection, nil
}



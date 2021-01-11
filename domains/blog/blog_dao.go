package blog

import (
	"context"
	"fmt"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/datasources/mongodb/blog_db"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO : return *errors_utils
func (blog *Blog) Get() error {
	db := blog_db.Client.Database("gRPC_Blog")
	collection := db.Collection("blogs")
	defer blog_db.Client.Disconnect(context.Background())

	res, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		// TODO Return Error
		logger_utils.Error("Error while Find Collection.", err)
		return err
	}

	logger_utils.Info(fmt.Sprintln(res))
	return nil

}

func (blog *Blog) Create() error {
	db := blog_db.Client.Database("gRPC_Blog")
	collection := db.Collection("blogs")

	res, err := collection.InsertOne(context.Background(), blog)
	if err != nil {
		logger_utils.Error("Error While Inserting Document", err)
		return err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		logger_utils.Error("Cannot convert to OID", err)
		return err
	}

	blog.ID = oid

	return nil

}

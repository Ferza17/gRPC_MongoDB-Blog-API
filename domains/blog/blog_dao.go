package blog

import (
	"context"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/datasources/mongodb/blog_db"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/errors_utils"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (blog *Blog) GetById() error {
	ctx := context.Background()
	if err := blog_db.Client.Connect(ctx); err != nil {
		return errors_utils.Internal("Unable to connect")
	}
	defer blog_db.Client.Disconnect(ctx)

	var result Blog
	filter := bson.D{primitive.E{Key: "_id", Value: blog.ID}}
	err := blog_db.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return errors_utils.Internal("Error to Find Data")
	}

	blog.AuthorId = result.AuthorId
	blog.Title = result.Title
	blog.Content = result.Content

	return nil

}

func (blog *Blog) Create() error {
	ctx := context.Background()
	if err := blog_db.Client.Connect(ctx); err != nil {
		return errors_utils.Internal("Unable to connect")
	}
	defer blog_db.Client.Disconnect(ctx)

	res, err := blog_db.Collection.InsertOne(context.Background(), blog)
	if err != nil {
		return errors_utils.Internal("Error While Inserting Document")
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return errors_utils.InvalidArgument("Invalid ID")
	}

	blog.ID = oid

	return nil

}

func (blog *Blog) Update() error {
	ctx := context.Background()
	if err := blog_db.Client.Connect(ctx); err != nil {
		return errors_utils.Internal("Unable to connect")
	}
	defer blog_db.Client.Disconnect(ctx)

	// FindById, if Doesnt exist then return error no ID match
	//var result Blog
	//filterId := bson.D{primitive.E{Key: "_id", Value: blog.ID}}
	//err := collection.FindOne(context.Background(), filterId).Decode(&result)
	//if err != nil {
	//	logger_utils.Error("Error while FindOne. ", err)
	//	return err
	//}

	// Update
	filter := bson.M{"_id": blog.ID}
	update := bson.D{
		{"$set", blog},
	}
	_, updateErr := blog_db.Collection.UpdateOne(context.TODO(), filter, update)

	if updateErr != nil {
		return errors_utils.Internal("Error While Update")
	}

	return nil

}

func (blog *Blog) Delete() error {
	ctx := context.Background()
	if err := blog_db.Client.Connect(ctx); err != nil {
		return errors_utils.Internal("Unable to connect")
	}
	defer blog_db.Client.Disconnect(ctx)

	filter := bson.M{"_id": blog.ID}

	res := blog_db.Collection.FindOneAndDelete(ctx, filter)

	if res.Err() != nil {
		return errors_utils.NotFound("ID Not Found")
	}

	return nil
}

func (blog *Blog) ListBlog() ([]Blog, error) {
	ctx := context.Background()
	if err := blog_db.Client.Connect(ctx); err != nil {
		return nil, errors_utils.Internal("Unable to connect")
	}
	defer blog_db.Client.Disconnect(ctx)

	cursor, err := blog_db.Collection.Find(ctx, bson.D{})
	if err != nil {
		logger_utils.Error("Error to Find Data ", err)
		return nil, errors_utils.Internal("Error to Find Data")
	}
	defer cursor.Close(ctx)

	var response []Blog

	for cursor.Next(ctx) {
		var data Blog
		if err := cursor.Decode(&data); err != nil {
			return nil, errors_utils.Internal("Error while Decode")
		}
		response = append(response, data)
	}

	return response, nil
}

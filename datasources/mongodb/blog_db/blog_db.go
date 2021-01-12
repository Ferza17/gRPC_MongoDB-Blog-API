package blog_db

import (
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/env_utils"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongodbUrl = env_utils.GetEnvironmentVariable("MONGODB_URL")
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
)

func init() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		logger_utils.Error("Error While Connect to MongoDB", err)
		return
	}

	Client = mongoClient
	Database = Client.Database("gRPC_Blog")
	Collection = Database.Collection("blogs")
	logger_utils.Info("Database Successfully configured")

}

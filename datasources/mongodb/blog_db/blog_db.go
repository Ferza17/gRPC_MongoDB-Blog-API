package blog_db

import (
	"context"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/env_utils"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mongodbUrl = env_utils.GetEnvironmentVariable("MONGODB_URL")
	Client     *mongo.Client
)

func init() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		logger_utils.Error("Error While Connect to MongoDB", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)

	// use Blog Database
	logger_utils.Info("Database Successfully Connected!")
	Client = mongoClient
}

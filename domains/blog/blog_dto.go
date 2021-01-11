package blog

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID       primitive.ObjectID `bson:"_id"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

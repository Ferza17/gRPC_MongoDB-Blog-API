package blog_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
)

var (
	BlogServices blogServiceInterface = &blogServiceStruct{}
)

type blogServiceStruct struct {
}

type blogServiceInterface interface {
	Get(id primitive.ObjectID) (*blog.Blog, error)
	Create(blog blog.Blog) (*blog.Blog, error)
}

func (s *blogServiceStruct) Get(id primitive.ObjectID) (*blog.Blog, error) {
	result := &blog.Blog{ID: id}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *blogServiceStruct) Create(blog blog.Blog) (*blog.Blog, error) {
	if err := blog.Create(); err != nil {
		return nil, err
	}

	return &blog, nil
}

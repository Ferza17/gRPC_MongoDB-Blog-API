package blog_service

import (
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/errors_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	BlogServices blogServiceInterface = &blogServiceStruct{}
)

type blogServiceStruct struct {
}

type blogServiceInterface interface {
	GetById(id primitive.ObjectID) (*blog.Blog, error)
	Create(blog blog.Blog) (*blog.Blog, error)
	Update(blog blog.Blog) (*blog.Blog, error)
	Delete(blog blog.Blog) (*blog.Blog, error)
	ListBlog(blog blog.Blog) ([]blog.Blog, error)
}

func (s *blogServiceStruct) GetById(id primitive.ObjectID) (*blog.Blog, error) {
	result := &blog.Blog{ID: id}

	if err := result.GetById(); err != nil {
		return nil, errors_utils.NotFound("ID Not Found")
	}
	return result, nil
}

func (s *blogServiceStruct) Create(blog blog.Blog) (*blog.Blog, error) {
	if err := blog.Create(); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (s *blogServiceStruct) Update(blog blog.Blog) (*blog.Blog, error) {
	if err := blog.Update(); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (s *blogServiceStruct) Delete(blog blog.Blog) (*blog.Blog, error) {
	if err := blog.Delete(); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (s *blogServiceStruct) ListBlog(blog blog.Blog) ([]blog.Blog, error) {
	res, err := blog.ListBlog()
	if err != nil {
		return nil, err
	}
	return res, nil
}

package blog_controller

import (
	"context"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/services/blog_service"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/logger_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
}

func (s *Server) CreateBlog(ctx context.Context, req *blog_proto.CreateBlogRequest) (*blog_proto.CreateBlogResponse, error) {

	res, err := blog_service.BlogServices.Create(blog.Blog{
		ID:       primitive.NewObjectID(),
		Title:    req.GetBlog().GetTitle(),
		Content:  req.GetBlog().GetContent(),
		AuthorId: req.GetBlog().GetAuthorId(),
	})

	if err != nil {
		return nil, err
	}

	return &blog_proto.CreateBlogResponse{
		Blog: &blog_proto.Blog{
			Id:       res.ID.String(),
			AuthorId: res.AuthorId,
			Title:    res.Title,
			Content:  res.Content,
		},
	}, nil

}

func (s *Server) ReadBlog(ctx context.Context, req *blog_proto.ReadBlogRequest) (*blog_proto.ReadBlogResponse, error) {
	blogId := req.GetBlogId()

	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		logger_utils.Error("Error While Convert ID", err)
		return nil, err
	}
	res, err := blog_service.BlogServices.GetById(oid)

	if res == nil {
		// TODO Not Found Error
		return nil, nil
	}

	return &blog_proto.ReadBlogResponse{
		Blog: &blog_proto.Blog{
			Id:       res.ID.String(),
			AuthorId: res.AuthorId,
			Title:    res.Title,
			Content:  res.Content,
		},
	}, nil
}

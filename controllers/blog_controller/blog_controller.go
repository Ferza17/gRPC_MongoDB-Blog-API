package blog_controller

import (
	"context"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/services/blog_service"
)

type Server struct {
}

//TODO: Implement with blog.pb.go

func (s *Server) CreateBlog(ctx context.Context, req *blog_proto.CreateBlogRequest) (*blog_proto.CreateBlogResponse, error) {

	res, err := blog_service.BlogServices.Create(blog.Blog{
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

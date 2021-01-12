package blog_controller

import (
	"context"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/services/blog_service"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/blog_utils"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/utils/errors_utils"
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
		Blog: blog_utils.DataToBlogPB(res),
	}, nil

}

func (s *Server) ReadBlog(ctx context.Context, req *blog_proto.ReadBlogRequest) (*blog_proto.ReadBlogResponse, error) {
	blogId := req.GetBlogId()

	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, errors_utils.InvalidArgument("Invalid ID")
	}
	res, err := blog_service.BlogServices.GetById(oid)
	if err != nil {
		return nil, errors_utils.NotFound("ID Not Found")
	}

	return &blog_proto.ReadBlogResponse{
		Blog: blog_utils.DataToBlogPB(res),
	}, nil
}

func (s *Server) UpdateBlog(ctx context.Context, req *blog_proto.UpdateBlogRequest) (*blog_proto.UpdateBlogResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetBlog().GetId())
	if err != nil {
		logger_utils.Error("Error While Convert ID", err)
		return nil, errors_utils.InvalidArgument("Invalid ID")
	}

	blogUpdate := blog.Blog{
		ID:       oid,
		AuthorId: req.GetBlog().GetAuthorId(),
		Title:    req.GetBlog().GetTitle(),
		Content:  req.GetBlog().GetContent(),
	}

	res, err := blog_service.BlogServices.Update(blogUpdate)

	return &blog_proto.UpdateBlogResponse{
		Blog: blog_utils.DataToBlogPB(res),
	}, nil
}

func (s *Server) DeleteBlog(ctx context.Context, req *blog_proto.DeleteBlogRequest) (*blog_proto.DeleteBLogResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		logger_utils.Error("Invalid ID", err)
		return nil, errors_utils.InvalidArgument("Invalid ID")
	}

	res, err := blog_service.BlogServices.Delete(blog.Blog{ID: oid})
	if err != nil {
		return nil, err
	}

	return &blog_proto.DeleteBLogResponse{
		BlogId: res.ID.String(),
	}, nil
}

func (s *Server) ListBlog(req *blog_proto.ListBlogRequest, stream blog_proto.BlogService_ListBlogServer) error {

	res, err := blog_service.BlogServices.ListBlog(blog.Blog{})
	if err != nil {
		return err
	}

	for _, item := range res {
		_ = stream.Send(&blog_proto.ListBlogResponse{Blog: blog_utils.DataToBlogPB(&item)})
	}
	// TODO : Streaming Client doesnt have any response

	return nil

}

package app

import (
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/controllers/blog_controller"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
	"google.golang.org/grpc"
)

func RpcAPI(s *grpc.Server)  {
	// Register gRPC
	blog_proto.RegisterBlogServiceServer(s, &blog_controller.Server{})
}

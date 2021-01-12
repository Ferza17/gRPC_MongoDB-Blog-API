package blog_utils

import (
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/domains/blog"
	"github.com/Ferza17/gRPC_MongoDB-Blog-API/protos/blog_proto"
)

func DataToBlogPB(data *blog.Blog) *blog_proto.Blog {
	return &blog_proto.Blog{
		Id:       data.ID.String(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

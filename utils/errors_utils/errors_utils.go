package errors_utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InvalidArgument(message string) error {
	return status.Error(codes.InvalidArgument, message)
}

func NotFound(message string) error {
	return status.Error(codes.NotFound, message)
}

func Internal(message string) error {
	return status.Error(codes.Internal, message)
}

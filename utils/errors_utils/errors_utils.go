package errors_utils

import (
	"google.golang.org/grpc/codes"
)

type RpcError struct {
	Code    codes.Code
	Message string
}

func Cancelled(message string) *RpcError {
	return &RpcError{
		Code:    codes.Canceled,
		Message: message,
	}
}

func InvalidArgument(message string) *RpcError {
	return &RpcError{
		Code: codes.InvalidArgument,
		Message: message,
	}
}

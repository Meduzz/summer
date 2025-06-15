package errors

import "github.com/Meduzz/summer/api"

func ParseError(err error) *api.RpcError {
	return &api.RpcError{
		Code:    ParseErrorCode,
		Message: err.Error(),
	}
}

func MethodNotFoundError() *api.RpcError {
	return &api.RpcError{
		Code:    MethodNotFoundCode,
		Message: "Method not found",
	}
}

func InternalError(err error) *api.RpcError {
	return &api.RpcError{
		Code:    InternalErrorCode,
		Message: err.Error(),
	}
}

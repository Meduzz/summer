package framework

import (
	"encoding/json"

	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
)

func ResultResponse(id json.RawMessage, result any) *api.Response {
	bs, err := json.Marshal(result)

	if err != nil {
		// TODO sharing this error is not very smart.
		return ErrorResponse(id, errors.InternalError(err))
	}

	return &api.Response{
		JsonRPC: api.JsonRPC,
		ID:      id,
		Result:  bs,
	}
}

func ErrorResponse(id json.RawMessage, err *api.RpcError) *api.Response {
	return &api.Response{
		JsonRPC: api.JsonRPC,
		ID:      id,
		Error:   err,
	}
}

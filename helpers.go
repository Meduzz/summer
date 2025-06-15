package summer

import (
	"encoding/json"

	"github.com/Meduzz/helper/http/client"
	"github.com/Meduzz/helper/http/herror"
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/Meduzz/summer/framework"
)

// TODO invest in a reflect based wrapper that can automagically translate params to func params?

// Wrap will wrap any (T)=>K, error func and allow you to turn it into a jsonrpc func.
func Wrap[T any, K any](delegate func(*T) (*K, error)) func(*api.Request) *api.Response {
	return func(r *api.Request) *api.Response {
		param := new(T)

		err := json.Unmarshal(r.Params, param)

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		resp, err := delegate(param)

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		return framework.ResultResponse(r.ID, resp)
	}
}

// Proxy will forward params to the provided url using the provided verb.
func HttpProxy(verb, url, contentType string) func(*api.Request) *api.Response {
	return func(r *api.Request) *api.Response {
		req, err := client.NewRequest(verb, url, r.Params, contentType)

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		res, err := req.DoDefault()

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		err = herror.ErrorFromCode(res.Code())

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		bs, err := res.AsBytes()

		if err != nil {
			return framework.ErrorResponse(r.ID, errors.InternalError(err))
		}

		return &api.Response{
			JsonRPC: api.JsonRPC,
			ID:      r.ID,
			Result:  bs,
		}
	}
}

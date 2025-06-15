package framework

import (
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/gin-gonic/gin"
)

type (
	Summer struct {
		handlers map[string]api.Handler
	}
)

func NewFramework() *Summer {
	return &Summer{
		handlers: make(map[string]api.Handler),
	}
}

func (s *Summer) Register(name string, handler api.Handler) {
	s.handlers[name] = handler
}

// TODO how to handle authentication authorization and pass it down the the delegates?
// HTTP this is a gin.HandlerFunc
func (s *Summer) HTTP(ctx *gin.Context) {
	// assume batching
	isBatch := true
	batchReq := api.BatchRequest{}
	err := ctx.ShouldBindBodyWithJSON(&batchReq)

	if err != nil {
		// falling back on single request
		req := &api.Request{}
		err = ctx.ShouldBindBodyWithJSON(req)

		if err != nil {
			// TODO did we just screw up not acting on first error?
			res := errors.ParseError(err)
			ctx.AbortWithStatusJSON(400, res)
			return
		}

		// but instantly pretend it was a batch
		isBatch = false
		batchReq = append(batchReq, req)
	}

	// request handler "loop"
	batchRes := slice.Map(batchReq, func(req *api.Request) *api.Response {
		// TODO how to make this async? (lots of extra work for non batch requests...)
		// 1. disallow nil returns. (requires notification and request handlers (requires more "infrastructure"))
		// 2. switch map to forEach and wrap each handler call in a go func(channel, req)
		// 3. iterate the channel for the expected number of responses then close it and return
		handler, ok := s.handlers[req.Method]

		if !ok {
			err := errors.MethodNotFoundError()
			return &api.Response{
				JsonRPC: api.JsonRPC,
				Error:   err,
				ID:      req.ID,
			}
		}

		return handler(req)
	})

	if batchRes == nil {
		ctx.Status(204)
		return
	}

	// remove notification responses
	batchRes = slice.Filter(batchRes, func(res *api.Response) bool {
		// TODO remove error responses without an id too?
		return res != nil
	})

	// "unbatch" if possible and requested
	if len(batchRes) == 1 && !isBatch {
		ctx.JSON(200, batchRes[0])
		return
	}

	ctx.JSON(200, batchRes)
}

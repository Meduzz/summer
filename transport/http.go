package transport

import (
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/Meduzz/summer/framework"
	"github.com/gin-gonic/gin"
)

// HTTP this is a gin.HandlerFunc
func HTTP(summer *framework.Summer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identity := framework.NewDummyIdentity()

		if summer.Auth != nil {
			identity = summer.Auth.Identity(ctx)
		}

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

			// instantly pretend it was a batch
			isBatch = false
			batchReq = append(batchReq, req)
		}

		// request handler "loop"
		batchRes := slice.Map(batchReq, func(req *api.Request) *api.Response {
			// TODO how to make this async? (lots of extra work for non batch requests...)
			// 1. disallow nil returns. (requires notification and request handlers (requires more "infrastructure"))
			// 2. switch map to forEach and wrap each handler call in a go func(channel, req)
			// 3. iterate the channel for the expected number of responses then close it and return

			// TODO is this enough, or will there still be cases where handlers need raw identity/roles?
			if identity.Method(req.Method) && identity.Request(req) {
				return summer.Handle(req)
			} else {
				// TODO invent error codes?
				return framework.ErrorResponse(req.ID, errors.MethodNotFoundError())
			}
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
}

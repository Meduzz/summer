package transport

import (
	"log"

	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/errors"
	"github.com/Meduzz/summer/framework"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func WS(summer *framework.Summer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identity := framework.NewDummyIdentity()

		if summer.Auth != nil {
			identity = summer.Auth.Identity(ctx)
		}

		ws := gin.WrapH(websocket.Handler(func(c *websocket.Conn) {
			for {
				req := &api.Request{}
				err := websocket.JSON.Receive(c, req)

				if err != nil {
					log.Println(err.Error()) // TODO pretty
					return
				}

				var res *api.Response
				if identity.Method(req.Method) && identity.Request(req) {
					res = summer.Handle(req)
				} else {
					res = framework.ErrorResponse(req.ID, errors.MethodNotFoundError())
				}

				if res != nil {
					err = websocket.JSON.Send(c, res)

					if err != nil {
						// TODO close connection if it's not dead already?
						log.Println(err.Error()) // TODO pretty
						return
					}
				}
			}
		}))

		ws(ctx)
	}
}

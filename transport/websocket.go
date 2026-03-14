package transport

import (
	"log"

	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/framework"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func WS(summer *framework.Summer) gin.HandlerFunc {
	return gin.WrapH(websocket.Handler(func(c *websocket.Conn) {
		for {
			req := &api.Request{}
			err := websocket.JSON.Receive(c, req)

			if err != nil {
				log.Println(err.Error()) // TODO pretty
				return
			}

			res := summer.Handle(req)

			if res != nil {
				err = websocket.JSON.Send(c, res)

				if err != nil {
					log.Println(err.Error()) // TODO pretty
					return
				}
			}
		}
	}))
}

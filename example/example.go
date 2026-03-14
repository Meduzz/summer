package main

import (
	"fmt"

	"github.com/Meduzz/disconnected"
	"github.com/Meduzz/disconnected/pkg/web"
	"github.com/Meduzz/summer"
	"github.com/gin-gonic/gin"
)

type (
	Greeting struct {
		Name string `json:"name"`
	}

	Result struct {
		Text string `json:"text"`
	}
)

func main() {
	summer.Register("greet", summer.Wrap(Greeter))
	summer.Register("proxy", summer.HttpProxy("POST", "http://localhost:8080/api/greet", "application/json"))

	disconnected.HttpServer("/", func(s *web.Server) {
		s.Static("/static", "static/")
		s.SPA("static/index.html")
		s.WithRouter(func(e *gin.Engine) {
			e.POST("/api/rpc", summer.HTTP())
			e.POST("/api/greet", func(ctx *gin.Context) {
				in := &Greeting{}
				ctx.BindJSON(in)

				out, _ := Greeter(in)
				ctx.JSON(200, out)
			})
			e.GET("/api/ws", summer.WS())
		})
	})
}

func Greeter(greeting *Greeting) (*Result, error) {
	return &Result{
		Text: fmt.Sprintf("Hello %s!", greeting.Name),
	}, nil
}

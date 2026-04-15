package main

import (
	"fmt"

	"github.com/Meduzz/disconnected"
	"github.com/Meduzz/disconnected/pkg/web"
	"github.com/Meduzz/helper/utilz"
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
	port := utilz.Env("PORT", "8080")
	summer.Register("greet", summer.Wrap(Greeter))
	summer.Register("proxy", summer.HttpProxy("POST", fmt.Sprintf("http://localhost:%s/api/greet", port), "application/json"))

	disconnected.HttpServer(func(s *web.Server) error {
		s.Static("/static", "static/")
		s.SPA("static/index.html")

		return s.WithRouter(func(e *gin.Engine) error {
			e.POST("/api/rpc", summer.HTTP())
			e.POST("/api/greet", func(ctx *gin.Context) {
				in := &Greeting{}
				ctx.BindJSON(in)

				out, _ := Greeter(in)
				ctx.JSON(200, out)
			})
			e.GET("/api/ws", summer.WS())
			return nil
		})
	})
}

func Greeter(greeting *Greeting) (*Result, error) {
	return &Result{
		Text: fmt.Sprintf("Hello %s!", greeting.Name),
	}, nil
}

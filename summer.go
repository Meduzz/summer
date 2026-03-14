package summer

import (
	"github.com/Meduzz/summer/api"
	"github.com/Meduzz/summer/framework"
	"github.com/Meduzz/summer/transport"
	"github.com/gin-gonic/gin"
)

// Register registers a handler to the provided name in the default framework instance
func Register(name string, handler api.Handler) {
	framework.Instance.Register(name, handler)
}

// HTTP is a gin.HandlerFunc based on the default framework instance
func HTTP() gin.HandlerFunc {
	return transport.HTTP(framework.Instance)
}

// Init creates a new instance of the framework
func Init() *framework.Summer {
	return framework.NewFramework()
}

func WS() gin.HandlerFunc {
	return transport.WS(framework.Instance)
}

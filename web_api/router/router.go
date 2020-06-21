package router

import (
	"github.com/gin-gonic/gin"

	"github.com/CapsLock-Studio/neo-komica/web_api/router/v1/foo"
)

// AddRouter - Add nested api routers to app.
func AddRouter(engine *gin.Engine) {
	r := engine.Group("/api/v1")

	r.GET("/foo", foo.Index)
}

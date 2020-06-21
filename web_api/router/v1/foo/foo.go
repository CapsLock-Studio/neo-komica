package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index - the test router for web_api
func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Test": "Hello World"})
}

package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorHandlerMiddleware is a middleware to cutsomize error responses.
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()

		var errors []*gin.Error

		if bindErrorToPrint := c.Errors.ByType(gin.ErrorTypeBind).Last(); bindErrorToPrint != nil {
			if err, ok := bindErrorToPrint.Err.(validator.ValidationErrors); ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"Errors": err,
				})
				return
			}
			// deal with other errors ...

			// Handle public errors.
		} else if errors = c.Errors.ByType(gin.ErrorTypePublic); len(errors) > 0 {

			c.JSON(http.StatusNotFound, gin.H{"Errors": getGinErrorMessages(errors)})

			// Handle private(internal) errors.
		} else if errors = c.Errors.ByType(gin.ErrorTypePrivate); len(errors) > 0 {
			response := gin.H{"Errors": getGinErrorMessages(errors)}
			debugStack := string(debug.Stack())

			if gin.Mode() != "release" {
				response = gin.H{"Errors": getGinErrorMessages(errors), "Stacktraces": []string{debugStack}}
			}

			c.JSON(
				http.StatusInternalServerError,
				response,
			)

			fmt.Println(debugStack)
		}
	}
}

func getGinErrorMessages(errors []*gin.Error) []string {
	messages := make([]string, 0, len(errors))

	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	return messages
}

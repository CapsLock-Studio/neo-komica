package logger

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

const rawDataKey = "rawDataKey"

// RequestLogger is the middleware for retain the request body for logger.
func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buf, _ := ctx.GetRawData()
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		ctx.Request.Body = rdr2

		ctx.Set(rawDataKey, rdr1)
		ctx.Next()
	}
}

// Error - log the error
func Error(ctx *gin.Context, err error) {
	ctx.Error(err)

	rawData, _ := ctx.Get(rawDataKey)
	reader := rawData.(io.Reader)

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetExtra("Request", gin.H{"Header": ctx.Request.Header, "Body": buf.String()})
		sentry.CaptureException(err)
	})
}

package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	_ "github.com/CapsLock-Studio/neo-komica/web_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/CapsLock-Studio/neo-komica/config"
	"github.com/CapsLock-Studio/neo-komica/model"
	"github.com/CapsLock-Studio/neo-komica/util"
	"github.com/CapsLock-Studio/neo-komica/web_api/middleware"
	"github.com/CapsLock-Studio/neo-komica/web_api/router"
	"github.com/CapsLock-Studio/neo-komica/web_api/util/logger"
)

// @title neo-komica golang API
// @version 0.1
// @description The API of neo-komica by golang

// @contact.name Calvin * Michael
// @contact.url https://youtu.be/PCp2iXA1uLE
// @contact.email calvin.huang@capslock.tw

// @BasePath /api/v1/
func main() {
	flag.Parse()

	defer glog.Flush()
	defer sentry.Flush(time.Second * 3)

	sentryConfig := config.NewSentry()

	if err := sentry.Init(sentry.ClientOptions{Dsn: sentryConfig.DSN}); err != nil {
		glog.Errorf("Sentry initialization failed: %v\n", err)
	}

	db := model.InitDB()
	defer db.Close()

	server := config.NewServer()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     server.AllowOrigins,
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "X-Accept-Language", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	app.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	app.Use(middleware.ErrorHandlerMiddleware())
	app.Use(logger.RequestLogger())

	app.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_DOC"))

	app.GET("/", func(ctx *gin.Context) { ctx.Status(http.StatusNoContent) })

	router.AddRouter(app)

	app.Run(util.Getenv("PORT", ":8080"))
}

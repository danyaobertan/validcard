package app

import (
	"fmt"
	"time"

	"net/http"

	validCard "github.com/danyaobertan/validcard/internal/api/delivery/http/validcard"
	validCardService "github.com/danyaobertan/validcard/internal/api/services/validcard"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 15 * time.Second
)

func (app *App) registerServices() {
	app.validCardService = validCardService.NewService()
}

func (app *App) registerHandlers() {
	app.validCardHTTPHandler = validCard.NewHandler(app.validCardService)
}

func (app *App) setupServerAndRoutes() error {
	router := gin.Default()
	app.registerRoutes(router)

	app.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout, // Helps protect against Slowloris attacks
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	return app.server.ListenAndServe()
}

func (app *App) registerRoutes(router *gin.Engine) {
	router.Static("/docs", "./docs")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.json")))

	router.POST("/validate", app.validCardHTTPHandler.ValidateCardInfo())
}

package app

import (
	"net/http"

	"github.com/danyaobertan/validcard/internal/api/delivery"
	"github.com/danyaobertan/validcard/internal/api/services"
)

type App struct {
	server *http.Server

	validCardService     services.ValidCard
	validCardHTTPHandler delivery.ValidCardHTTP
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {
	app.registerServices()
	app.registerHandlers()

	return app.setupServerAndRoutes()
}

func (app *App) Stop() error {
	return app.server.Close()
}

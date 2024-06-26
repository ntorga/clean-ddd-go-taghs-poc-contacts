package api

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	apiMiddleware "github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/api/middleware"
)

// @title			Clean DDD Full Stack Go PoC API
// @version			0.0.1
// @description		Clean Architecture & DDD with Go, Tailwind, Alpine.js, HTMX, and SQLite: A Proof of Concept

// @contact.name	Northon Torga
// @contact.url		https://ntorga.com/
// @contact.email	northontorga+github@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api
func ApiInit(e *echo.Echo, persistentDbSvc *db.PersistentDatabaseService) {
	basePath := "/api"
	baseRoute := e.Group(basePath)

	e.Pre(apiMiddleware.TrailingSlash(basePath))
	e.Use(apiMiddleware.PanicHandler)
	e.Use(apiMiddleware.SetDefaultHeaders(basePath))
	e.Use(apiMiddleware.SetDatabaseServices(persistentDbSvc))

	router := NewRouter(baseRoute, persistentDbSvc)
	router.RegisterRoutes()
}

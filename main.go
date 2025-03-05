package main

import (
	"fmt"
	"stationcontrol/cmd/routes"

	_ "stationcontrol/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           ListokControl
// @version         1.0
// @description     API Server for User Application
// @BasePath  /
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{fmt.Sprintf("http://%s:%s/swagger/doc.json", "127.0.0.1", "1323")}
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173",
			"http://10.0.0.4:4173",
			"http://10.0.0.4:5173",
			"http://localhost:4173"}, // Домен вашего фронтенда
		AllowMethods: []string{echo.GET, echo.POST},                              // Разрешённые методы
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization}, // Разрешённые заголовки
	}))

	apicontrol := e.Group("/control")

	routes.ControlRoutes(apicontrol)

	e.Logger.Fatal(e.Start(":1323"))
}

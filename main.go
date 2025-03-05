package main

import (
	"stationcontrol/cmd/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
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

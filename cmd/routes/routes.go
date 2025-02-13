package routes

import (
	"stationcontrol/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func ControlRoutes(group *echo.Group) {
	group.GET("/get-number", handlers.GetNumber)
	group.GET("/order-clean", handlers.OrderClean)
	group.GET("/print-test", handlers.PrintTest)
}

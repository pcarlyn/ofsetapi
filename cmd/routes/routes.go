package routes

import (
	"stationcontrol/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func ControlRoutes(group *echo.Group) {
	group.GET("/get-number", handlers.GetNumber)
	group.GET("/order-clean", handlers.OrderClean)
	group.GET("/print-test", handlers.PrintTest)
	group.GET("/get-lpf-info", handlers.GetLPFInfo)
	group.GET("/get-lpf-error", handlers.GetLPFError)
	group.GET("/run-speedtest", handlers.RunSpeedtest)
	group.GET("/run-speedtest-secure", handlers.RunSpeedtestSecure)
}

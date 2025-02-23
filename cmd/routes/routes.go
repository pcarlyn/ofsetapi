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
	group.GET("/get-log-prod/:date", handlers.GetLogProd)
	group.GET("/get-log-bot", handlers.GetLogBot)
	group.GET("/get-log-reboots", handlers.GetLogReboots)
	group.GET("/restart-bot", handlers.RestartBot)
	group.GET("/scanimage-test", handlers.RunScanerTest)
	group.GET("/check-usb-dir", handlers.CheckUSBDirectory)
	group.GET("/correct-usb-dir", handlers.CorrectUSBdir)
}

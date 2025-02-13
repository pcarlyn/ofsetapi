package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

func OrderClean(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	step1 := exec.Command("cancel", "-a")
	if err := step1.Run(); err != nil {
		fmt.Println(err)
	}
	step2 := exec.Command("cupsenable", "offset-pantum-lan")
	if err := step2.Run(); err != nil {
		fmt.Println(err)
	}
	step3 := exec.Command("systemctl", "restart", "cups")
	if err := step3.Run(); err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

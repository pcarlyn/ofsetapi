package handlers

import (
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

func RestartBot(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	cmd := exec.Command("systemctl", "restart", "offset@bot")
	if err := cmd.Run(); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute command: "+err.Error())
	}
	return c.String(http.StatusOK, "Success restart bot")
}

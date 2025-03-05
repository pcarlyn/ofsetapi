package handlers

import (
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"
	"time"

	"github.com/labstack/echo/v4"
)

// Handler for getting bot logs
// @Summary Получение логов бота
// @Description Возвращает логи сервиса @bot за последние 2 дня
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/get-log-bot [get]
func GetLogBot(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	date := time.Now().AddDate(0, 0, -2).Format("2006-01-02")

	cmd := exec.Command("journalctl", "-u", "offset@bot", "--since", date)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute command: "+err.Error())
	}

	return c.String(http.StatusOK, string(output))
}

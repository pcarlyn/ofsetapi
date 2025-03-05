package handlers

import (
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

// Handler for get the last 100 lines of LPF error log
// @Summary Получение последних 100 строк ошибок LPF
// @Description Возвращает последние 100 строк из файла лога ошибок LPF (`/opt/offset/log/lpf/error.log`).
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/get-lpf-error [get]
func GetLPFError(c echo.Context) error {

	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	cmd := exec.Command("tail", "-n", "100", "/opt/offset/log/lpf/error.log")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute command: "+err.Error())
	}

	return c.String(http.StatusOK, string(output))
}

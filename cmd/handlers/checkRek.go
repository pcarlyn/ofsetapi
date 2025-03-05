package handlers

import (
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

// Handler for get /var/www/reklama permissions
// @Summary Получение списка файлов и прав доступа в каталоге /var/www/reklama
// @Description Возвращает вывод команды `ls -lah` для каталога /var/www/reklama.
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/check-rek [get]
func GetLSRek(c echo.Context) error {

	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	cmd := exec.Command("ls", "-lah", "/var/www/reklama/")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute command: "+err.Error())
	}

	return c.String(http.StatusOK, string(output))
}

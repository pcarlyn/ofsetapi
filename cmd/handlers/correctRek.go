package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

// Handler for correcting /var/www/reklama permissions
// @Summary Исправление прав доступа для каталога /var/www/reklama
// @Description Устанавливает правильные права доступа и владельца для файлов в каталоге /var/www/reklama.
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/correct-rek [get]
func CorrectRek(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	step1 := exec.Command("sh", "-c", "chmod 755 /var/www/reklama/*")
	if err := step1.Run(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Ошибка при установке прав доступа")
	}

	step2 := exec.Command("sh", "-c", "chown www-data:www-data /var/www/reklama/*")
	if err := step2.Run(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Ошибка при изменении владельца")
	}

	return c.String(http.StatusOK, "Права исправлены")
}

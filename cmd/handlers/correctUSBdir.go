package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"stationcontrol/internal/config"
	"strings"

	"github.com/labstack/echo/v4"
)

// Handler for correcting USB directory
// @Summary Очистка каталога /media/user/
// @Description Проверяет существование каталога, проверяет его файловую систему (должна быть vfat) и удаляет все файлы в нём.
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/correct-usb-dir [get]
func CorrectUSBdir(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	if _, err := os.Stat("/media/user/"); os.IsNotExist(err) {
		return c.String(http.StatusBadRequest, "Каталог /media/user/ не найден")
	}

	cmd := exec.Command("df", "-T", "/media/user/")
	output, err := cmd.CombinedOutput()
	if err != nil || !strings.Contains(string(output), "vfat") {
		return c.String(http.StatusBadRequest, "Ошибка: /media/user/ не является флешкой")
	}

	step1 := exec.Command("find", "/media/user/", "-mindepth", "1", "-exec", "rm", "-rf", "{}", "+")

	if err := step1.Run(); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка удаления файлов: %v", err))
	}

	return c.String(http.StatusOK, "Файлы успешно удалены")
}

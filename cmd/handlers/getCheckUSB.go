package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

// Handler for checking USB directory and connected devices
// @Summary Проверка содержимого /media/user и списка USB-устройств
// @Description Выполняет команды `ls -lah /media/user` и `lsusb` для отображения содержимого USB-директории и списка подключённых USB-устройств.
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/check-usb-dir [get]
func CheckUSBDirectory(c echo.Context) error {

	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}
	cmd1 := exec.Command("ls", "-lah", "/media/user")
	output1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка выполнения ls: %v\n", err1))
	}

	cmd2 := exec.Command("lsusb")
	output2, err2 := cmd2.CombinedOutput()
	if err2 != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка выполнения lsusb: %v\n", err2))
	}

	output := fmt.Sprintf("Вывод `ls -lah /media/user`:\n%s\n\nВывод `lsusb`:\n%s", string(output1), string(output2))

	return c.String(http.StatusOK, output)
}

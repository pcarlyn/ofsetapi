package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"
	"time"

	"github.com/labstack/echo/v4"
)

// Handler for getting production logs
// @Summary Получение логов продакшена
// @Description Возвращает содержимое лог-файла за выбранную дату
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Param date path int true "Дата лога (1 - сегодня, 2 - вчера, 3 - позавчера)"
// @Router /control/get-log-prod/{date} [get]
func GetLogProd(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	dateID := c.Param("date")

	var date string

	switch dateID {
	case "1":
		date = time.Now().Format("02-01-06")
	case "2":
		date = time.Now().AddDate(0, 0, -1).Format("02-01-06")
	case "3":
		date = time.Now().AddDate(0, 0, -2).Format("02-01-06")
	default:
		return c.String(http.StatusBadRequest, "Invalid date ID")
	}

	filePath := fmt.Sprintf("/opt/listok/log/prod-log_%s.log", date)

	cmd := exec.Command("cat", filePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка выполнения команды:", err)
		return c.String(http.StatusInternalServerError, "Failed to execute command: "+err.Error())
	}
	return c.String(http.StatusOK, string(output))
}

package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"
	"time"

	"github.com/labstack/echo/v4"
)

// Handler for running speedtest
// @Summary Выполнение теста скорости интернета
// @Description Запускает тест скорости интернета с использованием команды `speedtest`. Возвращает результат или ошибку в зависимости от выполнения команды.
// @Tags control
// @Produce plain
// @Security ApiKeyAuth
// @Router /control/run-speedtest [get]
func RunSpeedtest(c echo.Context) error {

	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "speedtest")

	output, err := cmd.CombinedOutput()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return c.String(http.StatusRequestTimeout, "Speedtest timeout: выполнение команды превысило 1 минуту")
		}
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Speedtest error: %v\nOutput: %s", err, string(output)))
	}

	return c.String(http.StatusOK, string(output))
}

package handlers

import (
	"net/http"
	"os/exec"
	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

func GetLPFInfo(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	cmdTail := exec.Command("tail", "-n", "100", "/opt/offset/log/lpf/info.log")
	outputTail, err := cmdTail.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute tail: "+err.Error())
	}

	cmdLpq := exec.Command("lpq")
	outputLpq, err := cmdLpq.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to execute lpq: "+err.Error())
	}

	finalOutput := "=== LOG FILE ===\n" + string(outputTail) + "\n\n=== LPQ STATUS ===\n" + string(outputLpq)

	return c.String(http.StatusOK, finalOutput)
}

package handlers

import (
	"net/http"
	"os/exec"

	"stationcontrol/internal/config"

	"github.com/labstack/echo/v4"
)

func PrintTest(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != config.Cfg.TOKEN {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	cmd := exec.Command("lp")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get stdin pipe")
	}

	err = cmd.Start()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to start command")
	}

	_, err = stdin.Write([]byte("\n"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to write to stdin")
	}
	stdin.Close()

	err = cmd.Wait()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Command failed to finish")
	}

	return c.String(http.StatusOK, "Hello, World!")
}

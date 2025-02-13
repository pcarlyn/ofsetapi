package handlers

import (
	"net/http"
	"stationcontrol/internal/utils"

	"github.com/labstack/echo/v4"
)

func GetNumber(c echo.Context) error {
	num := utils.GetHostnameNumber()
	return c.String(http.StatusOK, num)

}

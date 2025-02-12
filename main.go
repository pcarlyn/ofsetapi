package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},                          // Домен вашего фронтенда
		AllowMethods: []string{echo.GET, echo.POST},                              // Разрешённые методы
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization}, // Разрешённые заголовки
	}))
	e.GET("/", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != "fdasdfsd" {
			return nil
		}
		step1 := exec.Command("cancel", "-a")
		if err := step1.Run(); err != nil {
			fmt.Println(err)
		}
		step2 := exec.Command("cupsenable", "offset-pantum-lan")
		if err := step2.Run(); err != nil {
			fmt.Println(err)
		}
		step3 := exec.Command("systemctl", "restart", "cups")
		if err := step3.Run(); err != nil {
			fmt.Println(err)
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

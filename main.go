package main

import (
	"fmt"
	"net/http"

	"github.com/GraphZC/demo-backend/config"
	"github.com/GraphZC/demo-backend/requests"
	"github.com/labstack/echo/v4"
)

func main() {
	configs := config.NewConfig()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Backend is running",
		})
	})

	e.POST("/secret", func(c echo.Context) error {
		var req requests.SecretRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}

		if req.Secret != configs.SecretPassword {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Invalid secret",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "The secret is not a secret",
		})
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", configs.Port)))
}

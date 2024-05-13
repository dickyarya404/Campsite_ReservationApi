package middleware

import (
	"campsite_reservation/utils"

	"github.com/labstack/echo/v4"
)

func Admin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get JwtToken from cookie
		JwtToken, err := c.Cookie("JwtToken")
		if err != nil {
			return c.JSON(401, map[string]interface{}{
				"message": "Unauthorized",
			})
		}

		// Decode JWT Token Payload
		payload, err := utils.DecodePayload(JwtToken.Value)
		if err != nil {
			return c.JSON(401, map[string]interface{}{
				"message": "Invalid JWT Token",
			})
		}

		role, ok := payload["role"].(string)
		if !ok || (role != "admin" && role != "super_admin") {
			return c.JSON(401, map[string]interface{}{
				"message": "Unauthorized",
			})
		}

		return next(c)
	}
}

package middlewares

import (
	"mini-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func UserAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userData := c.Locals("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		var user models.User

		err := models.DB.Where("id = ?", userID).First(&user).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
		}
		if user.ID != userID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		return c.Next()
	}
}

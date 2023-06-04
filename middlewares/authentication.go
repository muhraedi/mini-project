package middlewares

import (
	"mini-project/helpers"

	"github.com/gofiber/fiber/v2"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthentication",
				"message": err.Error(),
			})
		}
		c.Locals("userData", verifyToken)
		return c.Next()
	}
}

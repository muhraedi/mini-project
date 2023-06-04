package controllers

import (
	"mini-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	var user models.User
	userID := uint(userData["id"].(float64))
	if err := models.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	var user models.User
	userID := uint(userData["id"].(float64))

	err := models.DB.Model(&user).Where("id = ?", userID).Updates(&user).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   "Bad request",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to GET data",
		"errors":  nil,
		"data":    user,
	})
}

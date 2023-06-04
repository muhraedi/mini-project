package controllers

import (
	"mini-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTrxs(c *fiber.Ctx) error {
	var trx []models.Category
	models.DB.Find(&trx)

	return c.JSON(trx)
}

func GetTrxById(c *fiber.Ctx) error {
	id := c.Params("id")

	var trx models.Trx
	if err := models.DB.First(&trx, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(trx)
}

func PostTrx(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)

	var trx models.Trx
	userID := uint(userData["id"].(float64))

	trx.UserID = userID

	err := models.DB.Debug().Create(&trx).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(trx)
}

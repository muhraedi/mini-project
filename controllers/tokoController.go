package controllers

import (
	"mini-project/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyToko(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	var toko models.Toko
	userID := uint(userData["id"].(float64))

	if err := models.DB.First(&toko, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(toko)
}

func UpdateToko(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	var toko models.Toko

	id, _ := strconv.Atoi(c.Params("id"))
	userID := uint(userData["id"].(float64))

	if err := c.BodyParser(&toko); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	toko.UserID = userID
	toko.ID = uint(id)

	if models.DB.Model(&toko).Where("id = ?", id).Updates(&toko).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func GetAllToko(c *fiber.Ctx) error {
	var toko []models.Toko
	models.DB.Find(&toko)

	return c.JSON(toko)
}

func GetTokoById(c *fiber.Ctx) error {
	tokoId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
	}

	userData := c.Locals("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var toko models.Toko

	err = models.DB.Select("user_id").First(&toko, uint(tokoId)).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Data Not Found",
			"message": "data doesn't exist",
		})
	}

	if toko.UserID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "you are not allowed to access this data",
		})
	}

	return c.JSON(toko)
}

package controllers

import (
	"mini-project/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetAllAlamat(c *fiber.Ctx) error {
	var alamat []models.Alamat
	models.DB.Find(&alamat)

	return c.JSON(alamat)
}

func GetAlamatById(c *fiber.Ctx) error {
	alamatId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
	}

	userData := c.Locals("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var alamat models.Alamat

	err = models.DB.Select("user_id").First(&alamat, uint(alamatId)).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Data Not Found",
			"message": "data doesn't exist",
		})
	}

	if alamat.UserID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "you are not allowed to access this data",
		})
	}

	return c.JSON(alamat)
}

func PostAlamat(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	var alamat models.Alamat
	userID := uint(userData["id"].(float64))

	if err := c.BodyParser(&alamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	alamat.UserID = userID

	if err := models.DB.Debug().Create(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(alamat)
}

func UpdateAlamat(c *fiber.Ctx) error {
	userData := c.Locals("userData").(jwt.MapClaims)
	id, _ := strconv.Atoi(c.Params("id"))
	userID := uint(userData["id"].(float64))

	var alamat models.Alamat
	if err := c.BodyParser(&alamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	alamat.UserID = userID
	alamat.ID = uint(id)

	if models.DB.Model(&alamat).Where("id = ?", id).Updates(&alamat).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func DeleteAlamat(c *fiber.Ctx) error {
	id := c.Params("id")

	var alamat models.Alamat

	if models.DB.Delete(&alamat, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}

package controllers

import (
	"mini-project/helpers"
	"mini-project/models"

	"github.com/gofiber/fiber/v2"
)

func PostLogin(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	password := user.KataSandi

	err := models.DB.Debug().Where("no_telp = ?", user.NoTelp).Take(&user).Error

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  "No Telp salah",
			"data":    nil,
		})
	}

	comparePass := helpers.ComparePass([]byte(user.KataSandi), []byte(password))

	if !comparePass {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  "Kata sandi salah",
			"data":    nil,
		})
	}

	token := helpers.GenerateToken(user.ID, user.NoTelp)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data": fiber.Map{
			"nama":          user.Nama,
			"no_telp":       user.NoTelp,
			"tanggal_Lahir": user.TanggalLahir,
			"tentang":       user.Tentang,
			"pekerjaan":     user.Pekerjaan,
			"email":         user.Email,
			"id_provinsi":   user.IdProvinsi,
			"id_kota":       user.IdKota,
			"token":         token,
		},
	})
}

func PostRegister(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	toko := models.Toko{
		NamaToko: user.Nama + "Toko",
	}

	user.Tokos = toko

	if err := models.DB.Debug().Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST data",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to POST data",
		"errors":  nil,
		"data":    "Register Succeed",
	})
}

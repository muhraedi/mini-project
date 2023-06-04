package main

import (
	"mini-project/controllers"
	"mini-project/middlewares"
	"mini-project/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api/v1")

	// Category
	category := api.Group("/category")

	category.Use(middlewares.Authentication())
	category.Get("/", middlewares.UserAuthorization(), controllers.GetAllCategories)
	category.Get("/:id", middlewares.UserAuthorization(), controllers.GetCategoryById)
	category.Post("/", middlewares.UserAuthorization(), controllers.PostCategory)
	category.Put("/:id", middlewares.UserAuthorization(), controllers.UpdateCategory)
	category.Delete("/:id", middlewares.UserAuthorization(), controllers.DeleteCategory)

	// Auth
	auth := api.Group("/auth")

	auth.Post("/login", controllers.PostLogin)
	auth.Post("/register", controllers.PostRegister)

	// User
	user := api.Group("/user")

	user.Use(middlewares.Authentication())
	user.Get("/", middlewares.UserAuthorization(), controllers.GetUser)
	user.Put("/", middlewares.UserAuthorization(), controllers.UpdateUser)
	user.Get("/alamat", middlewares.UserAuthorization(), controllers.GetAllAlamat)
	user.Get("/alamat/:id", middlewares.UserAuthorization(), controllers.GetAlamatById)
	user.Post("/alamat/", middlewares.UserAuthorization(), controllers.PostAlamat)
	user.Put("/alamat/:id", middlewares.UserAuthorization(), controllers.UpdateAlamat)
	user.Delete("/alamat/:id", middlewares.UserAuthorization(), controllers.DeleteAlamat)

	// Toko
	toko := api.Group("/toko")

	toko.Use(middlewares.Authentication())
	toko.Get("/my", middlewares.UserAuthorization(), controllers.GetMyToko)
	toko.Put("/:id", middlewares.UserAuthorization(), controllers.UpdateToko)
	toko.Get("/:id", middlewares.UserAuthorization(), controllers.GetTokoById)
	toko.Get("/", middlewares.UserAuthorization(), controllers.GetAllToko)

	// Product
	product := api.Group("/product")

	product.Use(middlewares.Authentication())
	product.Get("/", middlewares.UserAuthorization(), controllers.GetAllProducts)
	product.Get("/:id", middlewares.UserAuthorization(), controllers.GetProductById)
	product.Post("/", middlewares.UserAuthorization(), controllers.PostProduct)
	product.Put("/:id", middlewares.UserAuthorization(), controllers.UpdateProduct)
	product.Delete("/:id", middlewares.UserAuthorization(), controllers.DeleteProduct)

	// Trx
	trx := api.Group("/trx")

	trx.Use(middlewares.Authentication())
	trx.Get("/", controllers.GetAllTrxs)
	trx.Get("/:id", controllers.GetTrxById)
	trx.Post("/", controllers.PostTrx)

	// Province City
	// provcity := api.Group("/provcity")

	// provcity.Get("/listprovincies")
	// provcity.Get("/listcities")
	// provcity.Get("/detailprovince/:prov_id")
	// provcity.Get("/detailcity/:city_id")

	app.Listen(":8000")
}

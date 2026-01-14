package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/controllers"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func UserRoutes(app *fiber.App, userController *controllers.UserController) { //ประกาศรับตัวแปร พร้อมกำหนด type ให้มันด้วย
	userGroup := app.Group("/api")
	userGroup.Post("/users", userController.CreateUser)
	userGroup.Get("/users", userController.GetUsers)
	userGroup.Get("/users/:id", userController.GetIdUsers)
	userGroup.Put("/users/:id", userController.UpdateUsers)
	userGroup.Delete("/users/:id", userController.DeleteUsers)
	userGroup.Get("/swagger/*", fiberSwagger.WrapHandler)
}

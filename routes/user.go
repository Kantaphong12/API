package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/controllers"
)

func UserRoutes(app *fiber.App, userController *controllers.UserController) { //ประกาศรับตัวแปร พร้อมกำหนด type ให้มันด้วย
	app.Post("/users", userController.CreateUser)
	app.Get("/users", userController.GetUsers)
	app.Get("/users/:id", userController.GetIdUsers)
	app.Put("/users/:id", userController.UpdateUsers)
	app.Delete("/users/:id", userController.DeleteUsers)
}

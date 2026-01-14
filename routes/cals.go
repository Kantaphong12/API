package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/controllers"
	"github.com/kantaphong/myapp/middleware"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func CalsRoutes(app *fiber.App, CalorieController *controllers.CalorieController) { // กำหนด route ของ user
	calsGroup := app.Group("/api") // สร้าง group ของ user
	// app.Get("/users/:id", CalorieController.GetIdUsers)
	calsGroup.Delete("/delete/:id", CalorieController.Delete)
	calsGroup.Post("/Register", CalorieController.Register)
	calsGroup.Post("/Login", CalorieController.Login)
	calsGroup.Get("/swagger/*", fiberSwagger.WrapHandler)

	protected := calsGroup.Use(middleware.AuthRequired) // ใช้ middleware ตรวจสอบโทเค็นก่อนเข้าถึง route ด้านล่าง

	protected.Post("/InsertCals", CalorieController.Create)
	protected.Put("/update/:id", CalorieController.Update)
	protected.Get("/GetCals/:id", CalorieController.GetData)
	protected.Put("/updateUserCals/:id", CalorieController.UpdateUsers)
}

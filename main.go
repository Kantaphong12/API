package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/controllers"
	"github.com/kantaphong/myapp/database"
	"github.com/kantaphong/myapp/routes"
	"github.com/kantaphong/myapp/services"
)

func main() {
	app := fiber.New()
	database.Connect()
	userService := services.NewUserService()                     //ได้ struct ที่มี function ที่สืบทอดตาม interface ต้องการ
	userController := controllers.NewUserController(userService) //ได้
	routes.UserRoutes(app, userController)                       //ส่ง app กับ userController ที่รวมทุกอย่างแล้วเข้าไปจัดการ route ที่ UserRoutes
	fmt.Println("Server running at http://localhost:5000")
	app.Listen(":5000")

}

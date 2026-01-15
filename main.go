// @title JWT CRUD API
// @version 1.0
// @description Basic CRUD with JWT Auth
// @host localhost:3000
// @BasePath /
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kantaphong/myapp/controllers"
	"github.com/kantaphong/myapp/database"
	"github.com/kantaphong/myapp/routes"
	"github.com/kantaphong/myapp/services"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "wildcard",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	database.Connect()
	userService := services.NewUserService()                     //ได้ struct ที่มี function ที่สืบทอดตาม interface ต้องการ
	userController := controllers.NewUserController(userService) //// {} คือสร้าง instance go ไม่อนุญาติให้ใช้ตรงๆ
	routes.UserRoutes(app, userController)

	calsService := services.NewCalsService()                        //ได้ struct ที่มี function ที่สืบทอดตาม interface ต้องการ
	calsController := controllers.NewCalorieController(calsService) //// {} คือสร้าง instance go ไม่อนุญาติให้ใช้ตรงๆ
	routes.CalsRoutes(app, calsController)                          //ส่ง app กับ userController ที่รวมทุกอย่างแล้วเข้าไปจัดการ route ที่ UserRoutes
	fmt.Println("Server running at http://localhost:5000")
	app.Listen(":5000")

}

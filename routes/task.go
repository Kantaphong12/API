package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/controllers"
	"github.com/kantaphong/myapp/middleware"
)

func TaskRoutes(app *fiber.App, taskController *controllers.TaskController) {
	taskGroup := app.Group("/taskapi") // สร้าง group ของ task
	protected := taskGroup.Use(middleware.AuthRequired)
	protected.Post("/insertTask", taskController.InsertTask) // สร้าง task ใหม่
	protected.Get("/getTask/:id", taskController.GetTask)    // ดึงข้อมูล task
	protected.Put("/updateTask/:id", taskController.Update)
}

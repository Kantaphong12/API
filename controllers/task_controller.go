package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/models"
	"github.com/kantaphong/myapp/services"
)

type TaskController struct {
	// You can add service dependencies here if needed
	servicesCurrent services.InterfaceTaskService
	Task            []models.Task
}

func NewTaskController(serviceSent services.InterfaceTaskService) *TaskController {
	return &TaskController{
		servicesCurrent: serviceSent,
	}
}

func (tc *TaskController) InsertTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if task.Title == "" { //ต้องมี title ทุกตัว
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}
	task, err := tc.servicesCurrent.InsertTask(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("เกิดข้อผิดพลาดในการสร้างข้อมูล")
	}

	return nil
}

func (tc *TaskController) GetTask(c *fiber.Ctx) error {
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	users, err := tc.servicesCurrent.GetTask(UserID)
	if err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.JSON(users)
}

func (tc *TaskController) Update(c *fiber.Ctx) error { // แก้ไขข้อมูล
	UserID, err := strconv.Atoi(c.Params("id")) // รับค่า ID จาก URL และแปลงเป็น int
	if err != nil {                             // ตรวจสอบข้อผิดพลาดในการแปลง ID จาก URL และแปลงเป็น int
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	} //ถ้าผ่านการตรวจสอบ
	var tasks models.Task                        // ประกาศตัวแปรรับข้อมูล
	if err := c.BodyParser(&tasks); err != nil { // แปลงข้อมูลที่รับมาเป็น struct
		println("แปรงข้อมูล", UserID)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // ถ้าแปลงไม่ได้ให้ส่ง error กลับ
	}
	updatedTasks, err := tc.servicesCurrent.Update(UserID, tasks) // เรียกใช้ service เพื่อแก้ไขข้อมูล
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("เกิดข้อผิดพลาดในการแก้ไขข้อมูล")
	}
	return c.JSON(updatedTasks) // ส่งข้อมูลที่แก้ไขแล้วกลับไป
}

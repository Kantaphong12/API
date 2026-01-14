package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/models"
	"github.com/kantaphong/myapp/services"
	"github.com/kantaphong/myapp/utils"
)

type CalorieController struct { // สร้าง struct ของ controller
	serviceCurrent services.InterfaceCalsService
	Cals           []models.Cals
}

func NewCalorieController(serviceSent services.InterfaceCalsService) *CalorieController { // สร้าง instance ของ controller
	return &CalorieController{
		serviceCurrent: serviceSent,
	}
}

func (uc *CalorieController) GetData(c *fiber.Ctx) error { // ดึงข้อมูลทั้งหมด
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	users, err := uc.serviceCurrent.GetAll(UserID)
	if err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.JSON(users)
}
func (uc *CalorieController) Create(c *fiber.Ctx) error { // สร้างข้อมูลใหม่

	var cals []models.Cals // ประกาศตัวแปรรับข้อมูล

	if err := c.BodyParser(&cals); err != nil { // แปลงข้อมูลที่รับมาเป็น struct
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // ถ้าแปลงไม่ได้ให้ส่ง error กลับ
	}

	if len(cals) == 0 { //check no data
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ //	ส่ง error กลับ
			"error": "No data provided",
		})

	}
	for _, cal := range cals { //validate data
		if cal.Foodname == "" { // ตรวจสอบว่า foodname ว่างหรือไม่
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Error": "Foodname is required",
			})
		}
	}

	cals, err := uc.serviceCurrent.Create(cals) // เรียกใช้ service เพื่อสร้างข้อมูล
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("เกิดข้อผิดพลาดในการสร้างข้อมูล")
	}
	return c.JSON(cals)
}
func (uc *CalorieController) Delete(c *fiber.Ctx) error {
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	_ = uc.serviceCurrent.Delete(UserID)

	return c.JSON(UserID)
}

func (uc *CalorieController) Register(c *fiber.Ctx) error {
	var userCals models.UserCals
	if err := c.BodyParser(&userCals); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	insertUser, err := uc.serviceCurrent.Register(userCals)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("เกิดข้อผิดพลาดในการลงทะเบียน")
	}
	return c.JSON(insertUser)
}

func (uc *CalorieController) Login(c *fiber.Ctx) error {
	var userCals models.UserCals
	if err := c.BodyParser(&userCals); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	getUser, err := uc.serviceCurrent.Login(userCals.Email, userCals.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("ข้อมูลผู้ใช้ไม่ถูกต้อง")
	}
	// const ID = 2
	token := utils.GenerateToken(getUser.ID)
	println(token)
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token, // 👈 ส่ง token นี้กลับไปเก็บใน LocalStorage ฝั่ง Vue
		"user": fiber.Map{
			"id":    getUser.ID,
			"name":  getUser.Name,
			"email": getUser.Email,
		},
	})
}

func (uc *CalorieController) UpdateUsers(c *fiber.Ctx) error {
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var UserCals models.UserCals
	if err := c.BodyParser(&UserCals); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, _ := uc.serviceCurrent.UpdateUsers(UserID, UserCals)
	return c.JSON(user)
}
func (uc *CalorieController) Update(c *fiber.Ctx) error { // แก้ไขข้อมูล
	UserID, err := strconv.Atoi(c.Params("id")) // รับค่า ID จาก URL และแปลงเป็น int
	if err != nil {                             // ตรวจสอบข้อผิดพลาดในการแปลง ID จาก URL และแปลงเป็น int
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	} //ถ้าผ่านการตรวจสอบ
	var cals models.Cals                        // ประกาศตัวแปรรับข้อมูล
	if err := c.BodyParser(&cals); err != nil { // แปลงข้อมูลที่รับมาเป็น struct
		println("แปรงข้อมูล", UserID)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // ถ้าแปลงไม่ได้ให้ส่ง error กลับ
	}
	updatedCals, err := uc.serviceCurrent.Update(UserID, cals) // เรียกใช้ service เพื่อแก้ไขข้อมูล
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("เกิดข้อผิดพลาดในการแก้ไขข้อมูล")
	}
	return c.JSON(updatedCals) // ส่งข้อมูลที่แก้ไขแล้วกลับไป
}

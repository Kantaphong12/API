package controllers

//controller ไฟล์ที่คุยกับ request เช่นเวลา client เรียกใช้งาน function ต่างๆ
//รับค่าจาก route
//เรียก service ไปทำงาน
//ส่งค่าที่ได้จาก service ไปให้ client ที่เรียกมาเช่นข้อมูลทั้งหมด
//***สำคัญตอนเรียกใช้เนี้ยมันก็ใช้ในนี้ทั้งหมดเพราะมันไปดึงโค้ดตัวอื่นๆมาต่อๆใน controller นี้หมดแล้วที่ต้องแยกกันเพราะจะได้แก้ง่ายไปแก้ที่อื่นมันก็ดึงมาเลยไม่ต้องหายาก
//****ที่ต้องเข้าใจน่าจะมีแค่วิธีการที่มันส่ง function จากที่อื่นมา

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kantaphong/myapp/models"
	"github.com/kantaphong/myapp/services"
)

type UserController struct {
	serviceCurrent services.InterfaceUserService //ประกาศ struct ไว้เฉยๆว่าคนที่สืบทอดตัวนี้ไปต้องมี service อะไร
	User           []models.User
}

// function นี้คือตัวเตรียมทุกอย่างให้พร้อมไปเอาทุกอย่างมารวมไว้ที่นี้เวลาเรียกใช้ก็เหมือนว่าโค๊ดรวมกันที่นี้แล้ว
func NewUserController(serviceSent services.InterfaceUserService) *UserController { //*UserController ค่าที่ return
	return &UserController{ //กำหนดค่าให้ service ที่ประกาศไว้ด้วย &UserController
		serviceCurrent: serviceSent, //ส่งค่าใหม่ที่มี interface ที่ imprement ครบพร้อมใช้ กลับไปให้คนที่เรียกใช้
	}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.serviceCurrent.GetAll()
	if err != nil {
		return c.Status(500).SendString("เกิดข้อผิดพลาด")
	}
	return c.JSON(users)
}
func (uc *UserController) GetIdUsers(c *fiber.Ctx) error {

	UserID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, _ := uc.serviceCurrent.GetByID(UserID)
	return c.JSON(user)
}

func (uc *UserController) UpdateUsers(c *fiber.Ctx) error {
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var User models.User
	if err := c.BodyParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, _ := uc.serviceCurrent.UpdateUsers(UserID, User)
	fmt.Println("UserID", UserID)

	return c.JSON(user)

}
func (uc *UserController) DeleteUsers(c *fiber.Ctx) error {
	UserID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	_ = uc.serviceCurrent.Delete(UserID)

	return c.JSON(UserID)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user []models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, err := uc.serviceCurrent.Create(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(user)
}

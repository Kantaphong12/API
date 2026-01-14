package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kantaphong/myapp/utils"
)

func AuthRequired(c *fiber.Ctx) error {
	// 1. ดึง Header ชื่อ Authorization ออกมา
	var tokenString string               // ประกาศตัวแปรเก็บโทเค็น
	authHeader := c.Get("Authorization") // ดึงค่า Authorization จาก Header
	// 2. เช็คว่าเป็น Bearer token หรือไม่
	if strings.HasPrefix(authHeader, "Bearer ") { // ตรวจสอบว่ามีคำว่า "Bearer " อยู่ใน Header หรือไม่
		tokenString = strings.TrimPrefix(authHeader, "Bearer ") // ตัดส่วนแรกคำว่า "Bearer " ออกเพื่อให้เหลือแค่โทเค็น
	} else {
		tokenString = authHeader // ถ้าไม่มีคำว่า "Bearer " ให้ใช้ค่าทั้งหมดเป็นโทเค็น
	}

	if c.Cookies("token") != "" { //เผืออนาคตใช้ cookie
		tokenString = c.Cookies("token") // ดึงโทเค็นจากคุกกี้ถ้าไม่มีใน Header
	}

	if tokenString == "" { // ถ้าไม่มีโทเค็นให้ส่ง error กลับ
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: No token provided",
			"token": authHeader,
		})
	}

	// 3. เรียก utils.ParseToken เพื่อตรวจสอบลายเซ็นและวันหมดอายุ
	token, err := utils.ParseToken(tokenString) // แปลงโทเค็นกลับเป็นข้อมูล
	if err != nil {                             // ถ้าแปลงไม่ได้หรือโทเค็นไม่ถูกต้องให้ส่ง error กลับ
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token ไม่ตรง หรือหมดอายุ l",
		})
	}
	if !token.Valid {
		fmt.Println("claims.UserID", tokenString)
	}
	// 4. ถ้าผ่าน ให้ดึงข้อมูลใน Token (Claims) ออกมา
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["UserID"].(float64) // ดึง UserID จาก Claims
		c.Locals("UserID", uint(userID))     // เก็บ UserID ไว้ใน Context ของ Fiber
		// 5. ผ่านด่านได้! ไปทำ Controller ถัดไป
		fmt.Println("token ให้ผ่าน claims.UserID ", userID)
		return c.Next() // ผ่านการตรวจสอบแล้วให้ไปทำงานต่อ
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized: Invalid token claims",
	})

}

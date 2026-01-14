package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) //ดึงค่าคีย์ลับจากตัวแปรสภาพแวดล้อม (Environment Variable)

func GetJwtSecret() []byte { // ฟังก์ชันเพื่อรับคีย์ลับ JWT
	if len(JwtSecret) == 0 {
		return []byte("mysecretkey1234") // คีย์ลับเริ่มต้นถ้าไม่มีการตั้งค่าในตัวแปรสภาพแวดล้อม
	}
	return JwtSecret
}

func GenerateToken(UserID int) string { // ฟังก์ชันเพื่อสร้างโทเค็น JWT

	claims := jwt.MapClaims{ // กำหนดข้อมูลโทเค็น (claims)
		"UserID": UserID,                                        // เพิ่ม UserID ลงในข้อมูลโทเค็น
		"exp":    jwt.TimeFunc().Add(24 * 3 * time.Hour).Unix(), // กำหนดเวลาหมดอายุของโทเค็น (7 วัน)
		"iat":    jwt.TimeFunc().Unix(),                         // เวลาที่โทเค็นถูกสร้าง
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // สร้างโทเค็นใหม่โดยใช้วิธีการเซ็นชื่อ HS256
	tokenString, err := token.SignedString(GetJwtSecret())
	if err != nil {
		panic(err) // หากเกิดข้อผิดพลาดในการเซ็นชื่อโทเค็น ให้หยุดการทำงานและแสดงข้อผิดพลาด
	}
	return tokenString // ส่งคืนโทเค็นในรูปแบบสตริง
}

func ParseToken(tokenString string) (*jwt.Token, error) { // ฟังก์ชันเพื่อแยกวิเคราะห์โทเค็น JWT
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// เช็คว่าวิธีเข้ารหัสตรงกันไหม (ป้องกันการปลอมแปลง Header)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return GetJwtSecret(), nil
	})
}

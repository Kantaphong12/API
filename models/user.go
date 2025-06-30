package models

// กำหนดโครงสร้างข้อมูลที่ตรงกับ table ใน DB
// struct ใช้ตอน query หรือ insert

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"size:100"`
	Password  string `gorm:"size:100"`
	CreatedAt time.Time
}

package models

// กำหนดโครงสร้างข้อมูลที่ตรงกับ table ใน DB
// struct ใช้ตอน query หรือ insert

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"size:100"`
	Password  string `gorm:"size:100"`
	CreatedAt time.Time
}
type Cals struct {
	ID          int      `gorm:"primaryKey"`
	User_id     int      `json:"user_id"`
	Date        string   `json:"date"`
	Foodname    string   `json:"foodname"`
	Calories    *float64 `json:"calories"`
	ServingSize *float64 `json:"serving_size"`
	Carbs       *float64 `json:"carbs"`
	Protein     *float64 `json:"protein"`
	Fat         *float64 `json:"fat"`
	Leucine     *float64 `json:"leucine"`
	Magnesium   *float64 `json:"magnesium"`
	Zinc        *float64 `json:"zinc"`
}
type UserCals struct {
	ID            int     `gorm:"primaryKey"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Age           int     `json:"age"`
	Weight        float64 `json:"weight"`         // สำคัญสำหรับคำนวณแคล
	Height        float64 `json:"height"`         // สำคัญสำหรับคำนวณแคล
	Fat           float64 `json:"fat"`            // เปอร์เซ็นต์ไขมันในร่างกาย
	Gender        string  `json:"gender"`         // male/female
	ActivityLevel float64 `json:"activity_level"` // 1.2, 1.55, etc.
	TargetCals    int     `json:"target_cals"`    // คำนวณเสร็จแล้วมาเก็บตรงนี้ หรือจะคำนวณสดก็ได้
}

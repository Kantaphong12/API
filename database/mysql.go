package database

// database/database.go = ไฟล์ที่จัดการเรื่องการเชื่อมต่อฐานข้อมูลทั้งหมด
// ✅ สร้าง connection
// ✅ ทำ AutoMigrate (ถ้าต้องสร้าง table)
// ✅ export ฟังก์ชันให้คนอื่นเรียกใช้งานได้

import (
	"fmt"

	"github.com/kantaphong/myapp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var dbCals *gorm.DB

func Connect() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "surindev_healthy:kantaphong@tcp(surindev.com)/surindev_healthy?charset=utf8&parseTime=True&loc=Local", // data source name
		// DSN:                       "root:@tcp(127.0.0.1:3306)/surindev_healthy?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connect Error:", err)
		panic("failed to connect database")
	} else {
		fmt.Println("DB Connect successfuly")
	}

	// ✅ สร้างตาราง User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println("AutoMigrate Error:", err)
	}

	ConnectCals()

}
func GetDB() *gorm.DB { //ใครเรียกใช้ return db type *gorm.DB ให้
	return db
}

func ConnectCals() {
	var err error
	dbCals, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "surindev_healthy:kantaphong@tcp(surindev.com)/surindev_healthy?charset=utf8&parseTime=True&loc=Local",
		// DSN:                       "root:@tcp(127.0.0.1:3306)/surindev_healthy?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("DB (Calories) Connect Error:", err)
		panic("failed to connect database: caloriesdb")
	} else {
		fmt.Println("DB (Calories) Connect successfuly")
	}

	// ✅ สร้างตาราง Cals
	if err := dbCals.AutoMigrate(&models.Cals{}); err != nil {
		fmt.Println("AutoMigrate Cals Error:", err)
	}
	if err := dbCals.AutoMigrate(&models.UserCals{}); err != nil {
		fmt.Println("AutoMigrate UserCals Error:", err)
	}

}

func GetDBCals() *gorm.DB { //ใครเรียกใช้ return db type *gorm.DB ให้
	return dbCals
}

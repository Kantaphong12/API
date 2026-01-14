package services

import (
	"github.com/kantaphong/myapp/database"
	"github.com/kantaphong/myapp/models"
)

type InterfaceCalsService interface { // กำหนด method ที่ service ต้องมี
	Create(DataCreate []models.Cals) ([]models.Cals, error)
	GetAll(ID int) ([]models.Cals, error)
	Delete(ID int) error
	Update(ID int, dataUpdate models.Cals) (models.Cals, error)
	Register(userCals models.UserCals) (models.UserCals, error)
	Login(email string, password string) (models.UserCals, error)
	UpdateUsers(ID int, DataUpdate models.UserCals) (models.UserCals, error)
}

func NewCalsService() InterfaceCalsService { // สร้าง instance ของ service
	return &structCalsService{}
}

type structCalsService struct{}

func (s *structCalsService) GetAll(ID int) ([]models.Cals, error) { // ดึงข้อมูลทั้งหมด
	var cals []models.Cals

	db := database.GetDBCals()
	if err := db.Where("user_id = ?", ID).Find(&cals).Error; err != nil {
		return nil, err
	}
	return cals, nil
}
func (s *structCalsService) Create(DataCreate []models.Cals) ([]models.Cals, error) { // สร้างข้อมูลใหม่
	db := database.GetDBCals()
	// var users []models.User

	db.Create(&DataCreate)
	return DataCreate, nil

}
func (s *structCalsService) Delete(ID int) error { // ลบข้อมูล
	db := database.GetDBCals()
	var cals models.Cals
	if err := db.Delete(&cals, ID).Error; err != nil {
		return err
	}
	return nil
}
func (s *structCalsService) Update(ID int, dataUpdate models.Cals) (models.Cals, error) { // แก้ไขข้อมูล
	db := database.GetDBCals()                        // ดึง connection ของ DB
	var cals models.Cals                              // สร้างตัวแปรเก็บข้อมูล user
	if err := db.First(&cals, ID).Error; err != nil { // ค้นหาข้อมูล user ตาม ID
		return models.Cals{}, err
	}
	cals.ID = ID
	cals.Foodname = dataUpdate.Foodname
	cals.Calories = dataUpdate.Calories
	cals.ServingSize = dataUpdate.ServingSize
	cals.Carbs = dataUpdate.Carbs
	cals.Protein = dataUpdate.Protein
	cals.Fat = dataUpdate.Fat
	cals.Leucine = dataUpdate.Leucine
	cals.Magnesium = dataUpdate.Magnesium
	cals.Zinc = dataUpdate.Zinc
	if err := db.Save(&cals).Error; err != nil { // บันทึกข้อมูลที่แก้ไขลง DB
		return models.Cals{}, err
	}

	return cals, nil
}
func (s *structCalsService) Register(userCals models.UserCals) (models.UserCals, error) {
	db := database.GetDBCals()
	db.Create(&userCals)
	return userCals, nil
}

func (s *structCalsService) Login(email string, password string) (models.UserCals, error) {
	db := database.GetDBCals()
	var userCals models.UserCals
	if err := db.Where("email = ? AND password = ?", email, password).First(&userCals).Error; err != nil {
		return models.UserCals{}, err
	}
	return userCals, nil
}
func (s *structCalsService) UpdateUsers(ID int, DataUpdate models.UserCals) (models.UserCals, error) {

	db := database.GetDBCals()
	var userCals models.UserCals
	if err := db.First(&userCals, ID).Error; err != nil { //
		return models.UserCals{}, err

	}
	userCals.ID = DataUpdate.ID
	userCals.Name = DataUpdate.Name
	userCals.Age = DataUpdate.Age
	userCals.Weight = DataUpdate.Weight
	userCals.Height = DataUpdate.Height
	userCals.Fat = DataUpdate.Fat
	userCals.ActivityLevel = DataUpdate.ActivityLevel
	if err := db.Save(&userCals).Error; err != nil {
		println("updateUsers2")
		return models.UserCals{}, err
	}
	return userCals, nil
}

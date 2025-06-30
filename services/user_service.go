package services

//service ชั้นที่ทำงานกับข้อมูลจริง เช่นอ่านจาก database
//ดึงข้อมูลจาก database
//เขียน logic ภายใน

import (
	"github.com/kantaphong/myapp/database"
	"github.com/kantaphong/myapp/models"
)

type InterfaceUserService interface {
	GetAll() ([]models.User, error) // บอกว่าต้องมีฟังก์ชัน GetAll()
	GetByID(ID int) ([]models.User, error)
	UpdateUsers(ID int, DataUpdate models.User) (models.User, error)
	Delete(ID int) error
	Create(DataCreate []models.User) ([]models.User, error)
}

func NewUserService() InterfaceUserService { //จะ return interface ได้ต้องมีสัญญาเช่น interface มี getall ก้ต้องสร้างให้ structuserService สืบทอด มีอะไรเพิ่มอีกก็ต้องสร้างอีก
	return &structuserService{} //ส่งกลับที่อยู่ struct ที่มีการสืบทอด GetAll กับ structuserService เรียบร้อยตามที่ interface ต้องการ
}

type structuserService struct{} //ประกาศไว้ให้สืบทอดเพื่อใช้ functon ที่ interface ต้องการ

func (s *structuserService) GetAll() ([]models.User, error) { //สืบทอด structuserService มาแล้ว
	var users []models.User

	// ใช้ database.GetDB() เพื่อดึง instance ของ GORM DB
	db := database.GetDB()

	// ดึงข้อมูลทั้งหมดจากตาราง users
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
func (s *structuserService) GetByID(ID int) ([]models.User, error) {
	var users []models.User
	// ใช้ database.GetDB() เพื่อดึง instance ของ GORM DB
	db := database.GetDB()
	if err := db.Find(&users, ID).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func (s *structuserService) UpdateUsers(ID int, DataUpdate models.User) (models.User, error) {
	db := database.GetDB()

	DataUpdate.ID = ID
	if err := db.Save(DataUpdate).Error; err != nil {
		return models.User{}, err
	}
	// var users []models.User
	return DataUpdate, nil

}

func (s *structuserService) Delete(ID int) error {
	db := database.GetDB()
	var users models.User
	if err := db.Delete(&users, ID).Error; err != nil {
		return err
	}
	return nil
}
func (s *structuserService) Create(DataCreate []models.User) ([]models.User, error) {
	db := database.GetDB()
	// var users []models.User
	db.Create(&DataCreate)
	return DataCreate, nil

}

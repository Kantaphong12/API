package services

import (
	"github.com/kantaphong/myapp/database"
	"github.com/kantaphong/myapp/models"
)

type InterfaceTaskService interface {
	// Define methods for task service here
	InsertTask(models.Task) (models.Task, error)
	GetTask(int) ([]models.Task, error)
	Update(ID int, dataUpdate models.Task) (models.Task, error)
}

func NewTaskService() InterfaceTaskService { //จะ return interface ได้ต้องมีสัญญาเช่น interface มี getall ก้ต้องสร้างให้ structtaskService สืบทอด มีอะไรเพิ่มอีกก็ต้องสร้างอีก
	return &structtaskService{} //ส่งกลับที่อยู่ struct ที่มีการสืบทอด GetAll กับ structtaskService เรียบร้อยตามที่ interface ต้องการ
}

type structtaskService struct{} //ประกาศไว้ให้สืบทอดเพื่อใช้ functon ที่ interface ต้องการ
func (s *structtaskService) InsertTask(DataCreate models.Task) (models.Task, error) {
	db := database.GetDBTask()
	db.Create(&DataCreate)
	return DataCreate, nil
}
func (s *structtaskService) GetTask(ID int) ([]models.Task, error) { // ดึงข้อมูลทั้งหมด
	var tasks []models.Task

	db := database.GetDBTask()
	if err := db.Where("user_id = ?", ID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
func (s *structtaskService) Update(ID int, dataUpdate models.Task) (models.Task, error) { // แก้ไขข้อมูล
	db := database.GetDBTask()                         // ดึง connection ของ DB
	var tasks models.Task                              // สร้างตัวแปรเก็บข้อมูล user
	if err := db.First(&tasks, ID).Error; err != nil { // ค้นหาข้อมูล user ตาม ID
		return models.Task{}, err
	}
	tasks.ID = ID
	tasks.Title = dataUpdate.Title
	tasks.Description = dataUpdate.Description
	tasks.Priority = dataUpdate.Priority
	tasks.DueDate = dataUpdate.DueDate
	tasks.Recurring = dataUpdate.Recurring
	tasks.RecurringFrequency = dataUpdate.RecurringFrequency
	tasks.Status = dataUpdate.Status
	tasks.Completed = dataUpdate.Completed
	tasks.DateCompleted = dataUpdate.DateCompleted
	println(dataUpdate.DateCompleted)
	// println(dataUpdate.Completed)
	if err := db.Save(&tasks).Error; err != nil { // บันทึกข้อมูลที่แก้ไขลง DB
		return models.Task{}, err
	}

	return tasks, nil
}

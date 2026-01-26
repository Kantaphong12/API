package models

type Task struct {
	ID                 int    `gorm:"primaryKey" json:"id"` // เพิ่ม json:"id" และแนะนำใช้ uint สำหรับ ID
	UserID             int    `json:"userId"`
	Date_create        string `gorm:"size:50" json:"date_create"`
	Title              string `gorm:"size:255" json:"title"`
	Description        string `gorm:"type:text" json:"description"`
	Priority           string `gorm:"size:50" json:"priority"`
	DueDate            string `gorm:"size:50" json:"dueDate"`
	Recurring          bool   `json:"recurring"`
	RecurringFrequency string `gorm:"size:50" json:"recurringFrequency"`
	Status             string `gorm:"size:50" json:"status"`
	Completed          bool   `json:"completed"`
	DateCompleted      string `gorm:"size:50" json:"date_completed"`
}

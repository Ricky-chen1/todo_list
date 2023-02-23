package models

import "gorm.io/gorm"

//建立对应model
type Task struct {
	gorm.Model
	User    User   `gorm:"ForeignKey:Uid"`
	Uid     uint   `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"type:longtext"`
	Status  int    `gorm:"default:-1"` //-1为未完成，1为已完成
	AddTime int
	EndTime int
}

func (task *Task) TableName() string {
	return "task"
}

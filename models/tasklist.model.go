package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Prefix     string `json:"Prefix"`
	CurrentNum int    `json:"CurrentNum"`
	Notice     string `json:"Notice"`
	Status     int    `json:"Status"`
	ProjectID  uint
}

func (t *Task) GetAllTaskList() (TaskList []Task, err error) {
	if err := db.Model(&t).Preload("Taskitem").Find(&TaskList).Error; err != nil {
		return nil, err
	}
	return
}

func (t *Task) CreateTaskList(c *gin.Context) (err error) {
	if err := c.BindJSON(&t); err != nil {
		return err
	}
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return
}

func (t *Task) DeteleTaskList(c *gin.Context) (err error) {
	ID := c.Query("ID")
	if err := db.Unscoped().Model(&t).Where("ID = ?", ID).Delete(&t).Error; err != nil {
		return err
	}
	return
}

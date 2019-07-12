package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Tasklist struct {
	gorm.Model
	Title     string
	Taskitems []Taskitem
}

func (t *Tasklist) GetAllTaskList() (TaskList []Tasklist, err error) {
	if err := db.Model(&t).Preload("Taskitem").Find(&t).Error; err != nil {
		return nil, err
	}
	return
}

func (t *Tasklist) CreateTaskList(c *gin.Context) (err error) {
	if err := c.BindJSON(&t); err != nil {
		return err
	}
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return
}

func (t *Tasklist) DeteleTaskList(c *gin.Context) (err error) {
	ID := c.Query("ID")
	if err := db.Unscoped().Model(&t).Where("ID = ?", ID).Delete(&t).Error; err != nil {
		return err
	}
	return
}

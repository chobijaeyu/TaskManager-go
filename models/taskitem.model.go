package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Taskitem struct {
	gorm.Model
	Prefix     string `json:"Prefix"`
	CurrentNum int    `json:"CurrentNum"`
	Notice     string `json:"Notice"`
	Status     int    `json:"Status"`
	TasklistID uint
}

func (t *Taskitem) UpdateTaskItem(c *gin.Context) (err error) {
	ID := c.Query("ID")
	Status := c.Query("Status")
	if err := db.Model(&t).Where("ID = ?", ID).Update("Status", Status).Error; err != nil {
		return err
	}
	return
}

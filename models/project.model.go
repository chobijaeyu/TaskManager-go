package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model

	Title    string
	Deadline time.Time
	Tasklist []*Task
}

func (p *Project) GetProjectDetail(c *gin.Context) (projectList []Project, err error) {
	ID := c.Param("ID")
	// if err := db.Model(&p).Preload("Tasklist.Taskitems").Preload("Tasklist").Find(&projectList).Error; err != nil {
	if err := db.Model(&p).Preload("Tasklist").Where("ID = ?", ID).First(&projectList).Error; err != nil {
		return nil, err
	}
	return
}

func (p *Project) GetAllProject() (projectList []Project, err error) {
	if err := db.Model(&p).Find(&projectList).Error; err != nil {
		return nil, err
	}
	return

}

func (p *Project) CreateProject(c *gin.Context) (err error) {
	if err := c.BindJSON(&p); err != nil {
		return err
	}
	if err := db.Create(&p).Error; err != nil {
		return err
	}
	return
}

func (p *Project) DeleteProject(c *gin.Context) (err error) {
	ID := c.Query(("ID"))
	if err := db.Unscoped().Model(&p).Where("ID = ?", ID).Delete(&p).Error; err != nil {
		return err
	}
	return
}

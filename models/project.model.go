package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model

	Prefix   string
	Deadline time.Time
	EndNum   int
	Tasklist []*Task
}

func (p *Project) GetProjectDetail(c *gin.Context) (projectList Project, err error) {
	ID := c.Param("ID")
	// if err := db.Model(&p).Preload("Tasklist.Taskitems").Preload("Tasklist").Find(&projectList).Error; err != nil {
	if err := db.Model(&p).Preload("Tasklist").Where("ID = ?", ID).First(&projectList).Error; err != nil {
		return Project{}, nil
	}
	return
}

func (p *Project) GetAllProject() (projectList []Project, err error) {
	if err := db.Model(&p).Find(&projectList).Error; err != nil {
		return nil, err
	}
	return

}

func (p *Project) CreateProject(c *gin.Context) (ID uint, err error) {
	if err := c.BindJSON(&p); err != nil {
		return 0, nil
	}
	if err := db.Create(&p).Error; err != nil {
		return 0, nil
	}
	ID = p.ID
	return
}

func (p *Project) DeleteProject(c *gin.Context) (err error) {
	ID := c.Param(("ID"))
	if err := db.Preload("Tasklist").Unscoped().Model(&p).Where("ID = ?", ID).Delete(&p).Error; err != nil {
		return err
	}
	return
}

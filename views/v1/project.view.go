package v1

import (
	"CardTaskGo/models"
	"CardTaskGo/pkg/e"

	"github.com/gin-gonic/gin"
)

type ProjectView struct{}

func (p *ProjectView) GetProjectDetail(c *gin.Context) {
	s := models.Project{}
	projectDetailData, err := s.GetProjectDetail(c)
	if err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"msg":  err.Error(),
		})
		return
	}
	statusCode := e.SUCCESS
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  e.GetMsg(statusCode),
		"Data": projectDetailData,
	})
}

func (p *ProjectView) GetAllProject(c *gin.Context) {
	s := models.Project{}
	projectListData, err := s.GetAllProject()
	if err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"msg":  err.Error(),
		})
		return
	}
	statusCode := e.SUCCESS
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  e.GetMsg(statusCode),
		"Data": projectListData,
	})
}

func (p *ProjectView) CreateProject(c *gin.Context) {
	s := models.Project{}
	ID, err := s.CreateProject(c)
	if err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"err":  err.Error(),
		})
		return
	}
	statusCode := e.Created
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  e.GetMsg(statusCode),
		"ID":   ID,
	})
}

func (p *ProjectView) DeleteProject(c *gin.Context) {
	s := models.Project{}
	err := s.DeleteProject(c)
	if err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"err":  err.Error(),
		})
		return
	}
	statusCode := e.SUCCESS
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  e.GetMsg(statusCode),
	})
}

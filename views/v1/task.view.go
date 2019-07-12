package v1

import (
	"CardTaskGo/models"
	"CardTaskGo/pkg/e"
	"github.com/gin-gonic/gin"
)

type TaskViews struct{}

func (t *TaskViews) GetTask(c *gin.Context) {
	var (
		s models.Tasklist
	)
	tasklistData ,err := s.GetAllTaskList()
	if err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"msg":  e.GetMsg(statusCode),
		})
		return
	}
	statusCode := e.SUCCESS
	c.JSON(statusCode, gin.H{
		"code":      statusCode,
		"msg":       e.GetMsg(statusCode),
		"Data": tasklistData,
	})


}

func (t *TaskViews) CreateTask	(c *gin.Context)  {
	s := models.Tasklist{}
	if err := s.CreateTaskList(c); err != nil {
		statusCode := e.ERROR
		c.JSON(statusCode, gin.H{
			"code": statusCode,
			"msg":  e.GetMsg(statusCode),
			"err":  err.Error(),
		})
		return
	}
	statusCode := e.Created
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  e.GetMsg(statusCode),
	})


}

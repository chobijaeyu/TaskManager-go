package routers

import (
	"CardTaskGo/pkg/e"
	v1 "CardTaskGo/views/v1"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowHeaders:     []string{"Authorization", "Access-Control-Allow-Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowOrigins:     []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	var (
		taskView    v1.TaskViews
		projectView v1.ProjectView
	)
	hub := e.NewHub()
	r.GET("/ws", func(c *gin.Context) {
		e.ServeWs(hub, c.Writer, c.Request)
	})
	goodsRouters := r.Group("/api/v1/task")
	{
		goodsRouters.GET("/:ID", taskView.GetTask)
		goodsRouters.POST("", taskView.CreateTask)
		goodsRouters.PUT("", taskView.UpdateTask)
		//	goodsRouters.DELETE("", taskViews.)
	}
	projectRouter := r.Group("/api/v1/project")
	{
		projectRouter.GET("/:ID", projectView.GetProjectDetail)
		projectRouter.GET("", projectView.GetAllProject)
		projectRouter.POST("", projectView.CreateProject)
		//	goodsRouters.PUT("", taskViews.)
		projectRouter.DELETE("/:ID", projectView.DeleteProject)
	}

	return r
}

package routers

import (
	"CardTaskGo/conf"
	v1 "CardTaskGo/views/v1"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
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
	gin.SetMode(conf.RUN_MODE)
	var taskView v1.TaskViews
	goodsRouters := r.Group("/api/v1/task")
	// goodsRouters.Use(middleware.MID_JWT())
	{
		goodsRouters.GET("", taskView.GetTask)
		goodsRouters.POST("", taskView.CreateTask)
	//	goodsRouters.PUT("", taskViews.)
	//	goodsRouters.DELETE("", taskViews.)
	}

	return r
}

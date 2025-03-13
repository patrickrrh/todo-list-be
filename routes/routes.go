package routes

import (
	controller "todo_list/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	taskRoutes := route.Group("/task")
	{
		taskRoutes.POST("/list", controller.GetTaskListByStatus)
		taskRoutes.POST("/create", controller.PostTask)
		taskRoutes.PUT("/update", controller.PutTask)
		taskRoutes.POST("/delete", controller.DeleteTask)
		taskRoutes.PUT("/status", controller.PutTaskStatus)
	}

	subtaskRoutes := route.Group("/subtask")
	{
		subtaskRoutes.POST("/create", controller.PostSubtask)
		subtaskRoutes.PUT("/update", controller.PutSubtask)
		subtaskRoutes.POST("/delete", controller.DeleteSubtask)
		subtaskRoutes.PUT("/status", controller.PutSubtaskStatus)
	}

	return route
}

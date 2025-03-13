package controller

import (
	"net/http"
	"todo_list/service"
	model "todo_list/types"

	"github.com/gin-gonic/gin"
)

func GetTaskListByStatus(c *gin.Context) {
	var taskStatus model.InputTaskStatusModel

	inputError := c.ShouldBindJSON(&taskStatus)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	taskList, err := service.GetTaskListByStatus(taskStatus.TaskStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    taskList,
		"message": "Success Get On Going Task List",
	})
}

func PostTask(c *gin.Context) {
	var task model.PostTaskModel

	inputError := c.ShouldBindJSON(&task)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PostTask(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Post Task",
	})
}

func PutTask(c *gin.Context) {
	var task model.PutTaskModel

	inputError := c.ShouldBindJSON(&task)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PutTask(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Put Task",
	})
}

func DeleteTask(c *gin.Context) {
	var taskId model.DeleteTaskModel

	inputError := c.ShouldBindJSON(&taskId)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.DeleteTask(taskId.TaskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Delete Task",
	})
}

func PutTaskStatus(c *gin.Context) {
	var task model.PutTaskStatusModel

	inputError := c.ShouldBindJSON(&task)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PutTaskStatus(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Put Task Status",
	})
}

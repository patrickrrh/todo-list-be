package controller

import (
	"net/http"
	"todo_list/service"
	model "todo_list/types"

	"github.com/gin-gonic/gin"
)

func PostSubtask(c *gin.Context) {
	var subtask model.PostSubtaskModel

	inputError := c.ShouldBindJSON(&subtask)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PostSubtask(subtask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Post Subtask",
	})
}

func PutSubtask(c *gin.Context) {
	var subtask model.PutSubtaskModel

	inputError := c.ShouldBindJSON(&subtask)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PutSubtask(subtask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Put Subtask",
	})
}

func DeleteSubtask(c *gin.Context) {
	var subtaskId model.DeleteSubtaskModel

	inputError := c.ShouldBindJSON(&subtaskId)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.DeleteSubtask(subtaskId.SubtaskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Delete Subtask",
	})
}

func PutSubtaskStatus(c *gin.Context) {
	var subtask model.PutSubtaskStatusModel

	inputError := c.ShouldBindJSON(&subtask)
	if inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputError.Error(),
		})
		return
	}

	err := service.PutSubtaskStatus(subtask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Put Subtask Status",
	})
}

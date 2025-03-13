package service

import (
	"log"
	"todo_list/repository"
	model "todo_list/types"
	"todo_list/utils"
)

func GetTaskListByStatus(taskStatus int) ([]model.TaskModel, error) {
	taskList, err := repository.GetTaskListByStatusDB(taskStatus)
	if err != nil {
		log.Println("[service][task_service.go][GetTaskListByStatus] errors in query", err.Error())
		return nil, err
	}

	for i := range taskList {
		taskDueTime, err := utils.ParseDueTime(taskList[i].TaskDueDate, taskList[i].TaskDueTime)
		if err != nil {
			return nil, err
		}

		if utils.IsOverdue(taskDueTime) && taskList[i].TaskStatus == 0 {
			taskList[i].TaskStatus = 2
		}

		subtask, err := repository.GetSubtaskListByTaskIdDB(taskList[i].TaskId)
		if err != nil {
			log.Println("[service][task_service.go][GetTaskListByStatus] errors in subtask query", err.Error())
			return nil, err
		}

		for j := range subtask {
			subtaskDueTime, err := utils.ParseDueTime(subtask[j].SubtaskDueDate, subtask[j].SubtaskDueTime)
			if err != nil {
				return nil, err
			}

			if utils.IsOverdue(subtaskDueTime) && subtask[j].SubtaskStatus == 0 {
				subtask[j].SubtaskStatus = 2
			}
		}

		taskList[i].SubtaskList = subtask
	}

	return taskList, nil
}

func PostTask(task model.PostTaskModel) error {
	err := repository.PostTaskDB(task)

	if err != nil {
		log.Println("[service][task_service.go][PostTask] errors in query", err.Error())
		return err
	}

	return nil
}

func PutTask(task model.PutTaskModel) error {
	_, err := repository.GetTaskByIdDB(task.TaskId)

	if err != nil {
		log.Println("[service][task_service.go][PutTask] errors in query", err.Error())
		return err
	}

	if task.TaskDueDate != nil && *task.TaskDueDate == "" {
		task.TaskDueDate = nil
	}

	if task.TaskDueTime != nil && *task.TaskDueTime == "" {
		task.TaskDueTime = nil
	}

	err = repository.PutTaskDB(task)

	if err != nil {
		log.Println("[service][task_service.go][PutTask] errors in query", err.Error())
		return err
	}

	return nil
}

func DeleteTask(taskId int) error {
	_, err := repository.GetTaskByIdDB(taskId)

	if err != nil {
		log.Println("[service][task_service.go][DeleteTask] errors in query", err.Error())
		return err
	}

	err = repository.DeleteTaskDB(taskId)

	if err != nil {
		log.Println("[service][task_service.go][DeleteTask] errors in query", err.Error())
		return err
	}

	return nil
}

func PutTaskStatus(task model.PutTaskStatusModel) error {
	_, err := repository.GetTaskByIdDB(task.TaskId)

	if err != nil {
		log.Println("[service][task_service.go][PutTaskStatus] errors in query", err.Error())
		return err
	}

	err = repository.PutTaskStatusDB(task)

	if err != nil {
		log.Println("[service][task_service.go][PutTaskStatus] errors in query", err.Error())
		return err
	}

	return nil
}

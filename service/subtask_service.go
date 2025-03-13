package service

import (
	"log"
	"todo_list/repository"
	model "todo_list/types"
)

func PostSubtask(subtask model.PostSubtaskModel) error {
	err := repository.PostSubtaskDB(subtask)

	if err != nil {
		log.Println("[service][subtask_service.go][PostSubtask] errors in query", err.Error())
		return err
	}

	return nil
}

func PutSubtask(subtask model.PutSubtaskModel) error {
	_, err := repository.GetSubtaskById(subtask.SubtaskId)

	if err != nil {
		log.Println("[service][subtask_service.go][PutSubtask] errors in query", err.Error())
		return err
	}

	if subtask.SubtaskDueDate != nil && *subtask.SubtaskDueDate == "" {
		subtask.SubtaskDueDate = nil
	}

	if subtask.SUbtaskDueTime != nil && *subtask.SUbtaskDueTime == "" {
		subtask.SUbtaskDueTime = nil
	}

	err = repository.PutSubtaskDB(subtask)

	if err != nil {
		log.Println("[service][subtask_service.go][PutSubtask] errors in query", err.Error())
		return err
	}

	return nil
}

func DeleteSubtask(subtaskId int) error {
	_, err := repository.GetSubtaskById(subtaskId)

	if err != nil {
		log.Println("[service][subtask_service.go][DeleteSubtask] errors in query", err.Error())
		return err
	}

	err = repository.DeleteSubtaskDB(subtaskId)

	if err != nil {
		log.Println("[service][subtask_service.go][DeleteSubtask] errors in query", err.Error())
		return err
	}

	return nil
}

func PutSubtaskStatus(subtask model.PutSubtaskStatusModel) error {
	_, err := repository.GetSubtaskById(subtask.SubtaskId)

	if err != nil {
		log.Println("[service][subtask_service.go][PutSubtaskStatus] errors in query", err.Error())
		return err
	}

	err = repository.PutSubtaskStatusDB(subtask)

	if err != nil {
		log.Println("[service][subtask_service.go][PutSubtaskStatus] errors in query", err.Error())
		return err
	}

	return nil
}

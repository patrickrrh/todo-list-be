package repository

import (
	"log"
	postgresql "todo_list/config/db/postgresql"
	model "todo_list/types"
)

func GetTaskListByStatusDB(taskStatus int) ([]model.TaskModel, error) {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		SELECT
			task_id,
			task_title,
			COALESCE(task_desc, '') AS task_desc,
			task_status,
			COALESCE(TO_CHAR(task_due_date, 'YYYY-MM-DD'), '') AS task_due_date,
			COALESCE(TO_CHAR(task_due_time, 'HH24:MI:SS'), '') AS task_due_time,
			created_date,
			COALESCE(TO_CHAR(updated_date, 'YYYY-MM-DD'), '') AS updated_date
		FROM
			task
		WHERE
			task_status = $1
		ORDER BY
			created_date DESC;
	`

	rows, err := db.Query(query, taskStatus)
	if err != nil {
		log.Println("[repository][task_repository.go][GetOnGOingTaskListDB] errors in query", err.Error())
		return nil, err
	}
	defer rows.Close()

	var taskList []model.TaskModel
	for rows.Next() {
		var task model.TaskModel
		err := rows.Scan(
			&task.TaskId,
			&task.TaskTitle,
			&task.TaskDesc,
			&task.TaskStatus,
			&task.TaskDueDate,
			&task.TaskDueTime,
			&task.CreatedDate,
			&task.UpdatedDate,
		)
		if err != nil {
			log.Println("[repository][task_repository.go][GetOnGOingTaskListDB] errors in scan", err.Error())
			return nil, err
		}

		taskList = append(taskList, task)
	}

	return taskList, nil
}

func PostTaskDB(task model.PostTaskModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		INSERT INTO
			task
		(
			task_title,
			task_desc,
			task_due_date,
			task_due_time
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4
		);
	`

	_, err := db.Exec(query, task.TaskTitle, task.TaskDesc, task.TaskDueDate, task.TaskDueTime)
	if err != nil {
		log.Println("[repository][task_repository.go][PostTaskDB] errors in query", err.Error())
		return err
	}

	return nil
}

func GetTaskByIdDB(taskId int) (model.TaskModel, error) {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		SELECT
			task_id,
			task_title,
			COALESCE(task_desc, '') AS task_desc,
			task_status,
			COALESCE(TO_CHAR(task_due_date, 'YYYY-MM-DD'), '') AS task_due_date,
			COALESCE(TO_CHAR(task_due_time, 'HH24:MI:SS'), '') AS task_due_time,
			created_date,
			COALESCE(TO_CHAR(updated_date, 'YYYY-MM-DD'), '') AS updated_date
		FROM
			task
		WHERE
			task_id = $1;
	`

	var task model.TaskModel
	err := db.QueryRow(query, taskId).Scan(
		&task.TaskId,
		&task.TaskTitle,
		&task.TaskDesc,
		&task.TaskStatus,
		&task.TaskDueDate,
		&task.TaskDueTime,
		&task.CreatedDate,
		&task.UpdatedDate,
	)

	if err != nil {
		log.Println("[repository][task_repository.go][GetTaskByIdDB] errors in query", err.Error())
		return model.TaskModel{}, err
	}

	return task, nil
}

func PutTaskDB(task model.PutTaskModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		UPDATE
			task
		SET
			task_title = $1,
			task_desc = $2,
			task_due_date = $3,
			task_due_time = $4,
			updated_date = now()
		WHERE
			task_id = $5;
	`

	_, err := db.Exec(query, task.TaskTitle, task.TaskDesc, task.TaskDueDate, task.TaskDueTime, task.TaskId)
	if err != nil {
		log.Println("[repository][task_repository.go][PutTaskDB] errors in query", err.Error())
		return err
	}

	return nil
}

func DeleteTaskDB(taskId int) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println("[repository][task_repository.go][DeleteTaskDB] error starting transaction", err.Error())
		return err
	}

	deleteSubtaskQuery := `DELETE FROM subtask WHERE task_id = $1;`
	_, err = tx.Exec(deleteSubtaskQuery, taskId)
	if err != nil {
		tx.Rollback()
		log.Println("[repository][task_repository.go][DeleteTaskDB] error deleting subtasks", err.Error())
		return err
	}

	deleteTaskQuery := `DELETE FROM task WHERE task_id = $1;`
	_, err = tx.Exec(deleteTaskQuery, taskId)
	if err != nil {
		tx.Rollback()
		log.Println("[repository][task_repository.go][DeleteTaskDB] error deleting task", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("[repository][task_repository.go][DeleteTaskDB] error committing transaction", err.Error())
		return err
	}

	return nil
}

func PutTaskStatusDB(task model.PutTaskStatusModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		UPDATE
			task
		SET
			task_status = $1
		WHERE
			task_id = $2;
	`

	_, err := db.Exec(query, task.TaskStatus, task.TaskId)
	if err != nil {
		log.Println("[repository][task_repository.go][PutTaskStatusDB] errors in query", err.Error())
		return err
	}

	return nil
}

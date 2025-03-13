package repository

import (
	"log"
	postgresql "todo_list/config/db/postgresql"
	model "todo_list/types"
)

func GetSubtaskListByTaskIdDB(taskId int) ([]model.SubtaskModel, error) {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		SELECT
			subtask_id,
			task_id,
			subtask_title,
			COALESCE(subtask_desc, '') AS subtask_desc,
			subtask_status,
			COALESCE(TO_CHAR(subtask_due_date, 'YYYY-MM-DD'), '') AS subtask_due_date,
			COALESCE(TO_CHAR(subtask_due_time, 'HH24:MI:SS'), '') AS subtask_due_time,
			created_date,
			COALESCE(TO_CHAR(updated_date, 'YYYY-MM-DD'), '') AS updated_date
		FROM
			subtask
		WHERE
			task_id = $1
		ORDER BY
			created_date DESC
		;
	`

	rows, err := db.Query(query, taskId)
	if err != nil {
		log.Println("[repository][subtask_repository.go][GetSubtaskListByTaskIdDB] errors in query", err.Error())
		return nil, err
	}
	defer rows.Close()

	var subtaskList []model.SubtaskModel
	for rows.Next() {
		var subtask model.SubtaskModel
		err := rows.Scan(
			&subtask.SubtaskId,
			&subtask.TaskId,
			&subtask.SubtaskTitle,
			&subtask.SubtaskDesc,
			&subtask.SubtaskStatus,
			&subtask.SubtaskDueDate,
			&subtask.SubtaskDueTime,
			&subtask.CreatedDate,
			&subtask.UpdatedDate,
		)
		if err != nil {
			log.Println("[repository][subtask_repository.go][GetSubtaskListByTaskIdDB] errors in scan", err.Error())
			return nil, err
		}

		subtaskList = append(subtaskList, subtask)
	}

	return subtaskList, nil
}

func PostSubtaskDB(subtask model.PostSubtaskModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		INSERT INTO
			subtask
		(
			task_id,
			subtask_title,
			subtask_desc,
			subtask_due_date,
			subtask_due_time
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5
		);
	`

	_, err := db.Exec(query, subtask.TaskId, subtask.SubtaskTitle, subtask.SubtaskDesc, subtask.SubtaskDueDate, subtask.SUbtaskDueTime)
	if err != nil {
		log.Println("[repository][subtask_repository.go][PostSubtaskDB] errors in query", err.Error())
		return err
	}

	return nil
}

func GetSubtaskById(subtaskId int) (model.SubtaskModel, error) {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		SELECT
			subtask_id,
			task_id,
			subtask_title,
			COALESCE(subtask_desc, '') AS subtask_desc,
			subtask_status,
			COALESCE(TO_CHAR(subtask_due_date, 'YYYY-MM-DD'), '') AS subtask_due_date,
			COALESCE(TO_CHAR(subtask_due_time, 'HH24:MI:SS'), '') AS subtask_due_time,
			created_date,
			COALESCE(TO_CHAR(updated_date, 'YYYY-MM-DD'), '') AS updated_date
		FROM
			subtask
		WHERE
			subtask_id = $1;
	`

	var subtask model.SubtaskModel
	err := db.QueryRow(query, subtaskId).Scan(
		&subtask.SubtaskId,
		&subtask.TaskId,
		&subtask.SubtaskTitle,
		&subtask.SubtaskDesc,
		&subtask.SubtaskStatus,
		&subtask.SubtaskDueDate,
		&subtask.SubtaskDueTime,
		&subtask.CreatedDate,
		&subtask.UpdatedDate,
	)
	if err != nil {
		log.Println("[repository][subtask_repository.go][GetSubtaskById] errors in query", err.Error())
		return model.SubtaskModel{}, err
	}

	return subtask, nil
}

func PutSubtaskDB(subtask model.PutSubtaskModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		UPDATE
			subtask
		SET
			subtask_title = $1,
			subtask_desc = $2,
			subtask_due_date = $3,
			subtask_due_time = $4,
			updated_date = now()
		WHERE
			subtask_id = $5;
	`

	_, err := db.Exec(query, subtask.SubtaskTitle, subtask.SubtaskDesc, subtask.SubtaskDueDate, subtask.SUbtaskDueTime, subtask.SubtaskId)
	if err != nil {
		log.Println("[repository][subtask_repository.go][PutSubtaskDB] errors in query", err.Error())
		return err
	}

	return nil
}

func DeleteSubtaskDB(subtaskId int) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		DELETE
		FROM
			subtask
		WHERE
			subtask_id = $1;
	`

	_, err := db.Exec(query, subtaskId)
	if err != nil {
		log.Println("[repository][subtask_repository.go][DeleteSubtaskDB] errors in query", err.Error())
		return err
	}

	return nil
}

func PutSubtaskStatusDB(subtask model.PutSubtaskStatusModel) error {
	db := postgresql.OpenConnection()
	defer db.Close()

	query := `
		UPDATE
			subtask
		SET
			subtask_status = $1
		WHERE
			subtask_id = $2;
	`

	_, err := db.Exec(query, subtask.SubtaskStatus, subtask.SubtaskId)
	if err != nil {
		log.Println("[repository][subtask_repository.go][PutSubtaskStatusDB] errors in query", err.Error())
		return err
	}

	return nil
}

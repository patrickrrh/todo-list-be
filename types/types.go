package model

type TaskModel struct {
	TaskId      int            `json:"task_id"`
	TaskTitle   string         `json:"task_title"`
	TaskDesc    string         `json:"task_desc"`
	TaskStatus  int            `json:"task_status"`
	TaskDueDate string         `json:"task_due_date"`
	TaskDueTime string         `json:"task_due_time"`
	CreatedDate string         `json:"created_date"`
	UpdatedDate string         `json:"updated_date"`
	SubtaskList []SubtaskModel `json:"subtask_list"`
}

type InputTaskStatusModel struct {
	TaskStatus int `json:"task_status"`
}

type PostTaskModel struct {
	TaskTitle   string  `json:"task_title"`
	TaskDesc    *string `json:"task_desc"`
	TaskDueDate *string `json:"task_due_date"`
	TaskDueTime *string `json:"task_due_time"`
}

type PutTaskModel struct {
	TaskId      int     `json:"task_id"`
	TaskTitle   string  `json:"task_title"`
	TaskDesc    *string `json:"task_desc"`
	TaskDueDate *string `json:"task_due_date"`
	TaskDueTime *string `json:"task_due_time"`
	UpdatedDate string  `json:"updated_date"`
}

type DeleteTaskModel struct {
	TaskId int `json:"task_id"`
}

type PutTaskStatusModel struct {
	TaskId     int `json:"task_id"`
	TaskStatus int `json:"task_status"`
}

type SubtaskModel struct {
	SubtaskId      int    `json:"subtask_id"`
	TaskId         int    `json:"task_id"`
	SubtaskTitle   string `json:"subtask_title"`
	SubtaskDesc    string `json:"subtask_desc"`
	SubtaskStatus  int    `json:"subtask_status"`
	SubtaskDueDate string `json:"subtask_due_date"`
	SubtaskDueTime string `json:"subtask_due_time"`
	CreatedDate    string `json:"created_date"`
	UpdatedDate    string `json:"updated_date"`
}

type PostSubtaskModel struct {
	TaskId         int     `json:"task_id"`
	SubtaskTitle   string  `json:"subtask_title"`
	SubtaskDesc    *string `json:"subtask_desc"`
	SubtaskDueDate *string `json:"subtask_due_date"`
	SUbtaskDueTime *string `json:"subtask_due_time"`
}

type PutSubtaskModel struct {
	SubtaskId      int     `json:"subtask_id"`
	SubtaskTitle   string  `json:"subtask_title"`
	SubtaskDesc    *string `json:"subtask_desc"`
	SubtaskDueDate *string `json:"subtask_due_date"`
	SUbtaskDueTime *string `json:"subtask_due_time"`
	UpdatedDate    string  `json:"updated_date"`
}

type DeleteSubtaskModel struct {
	SubtaskId int `json:"subtask_id"`
}

type PutSubtaskStatusModel struct {
	SubtaskId     int `json:"subtask_id"`
	SubtaskStatus int `json:"subtask_status"`
}

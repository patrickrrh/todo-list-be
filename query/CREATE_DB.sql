CREATE TABLE public.task (
    task_id SERIAL PRIMARY KEY,
    task_title VARCHAR(255) NOT NULL,
    task_desc VARCHAR(1000),
    task_status INT DEFAULT 0,
    task_due_date DATE,
    task_due_time TIME,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP
);

CREATE TABLE public.subtask (
    subtask_id SERIAL PRIMARY KEY,
    task_id INT,
    FOREIGN KEY (task_id) REFERENCES public.task(task_id),
    subtask_title VARCHAR(255) NOT NULL,
    subtask_desc VARCHAR(1000),
    subtask_status INT DEFAULT 0,
    subtask_due_date DATE,
    subtask_due_time TIME,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP
);
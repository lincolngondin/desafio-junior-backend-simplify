package task

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewReposity(database *sql.DB) *repository {
	return &repository{database}
}

func (repo *repository) CreateTask(task *Task) error {
	_, err := repo.db.Exec("INSERT INTO tasks (nome, descricao, completado, prioridade) VALUES ($1, $2, $3, $4);",
		task.Name,
		task.Description,
		task.Completed,
		task.Priority,
	)
	return err
}

func (repo *repository) DeleteTask(id int64) error {
	_, err := repo.db.Exec("DELETE FROM tasks WHERE id = $1;", id)
	return err
}

func (repo *repository) GetTasks() ([]Task, error) {
	rows, err := repo.db.Query("SELECT * FROM tasks;")
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0, 10)
	for rows.Next() {
		task := NewDefaultTask()
		scanErr := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Completed, &task.Priority)
		if scanErr != nil {
			return nil, scanErr
		}
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

func (repo *repository) UpdateTask(id int64, name string, description string, priority Priority, completed bool) error {
    _, err := repo.db.Exec("UPDATE tasks SET nome = $1, descricao = $2, completado = $3, prioridade = $4 WHERE id = $5;",
		name,
		description,
		priority,
		completed,
		id,
	)
	return err
}

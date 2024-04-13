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
	return nil
}

func (repo *repository) GetTasks() []Task {
	return nil
}

func (repo *repository) UpdateTask(id int64, name string, description string, priority Priority, completed bool) error {
	return nil
}

package task

import (
	"database/sql"
	"errors"
)

var ErrorTaskNotFound = errors.New("Task Not Found")

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
    tx, txErr := repo.db.Begin()
    if txErr != nil {
        return txErr
    }
    row := tx.QueryRow("SELECT * FROM tasks WHERE id = $1;", id)
    taskExample := NewDefaultTask()
    noRows := row.Scan(&taskExample.Id, &taskExample.Name, &taskExample.Description, &taskExample.Completed, &taskExample.Priority)
    if noRows == sql.ErrNoRows {
        tx.Rollback()
        return ErrorTaskNotFound
    }
    if noRows != nil {
        tx.Rollback()
        return noRows
    }
	_, err := tx.Exec("DELETE FROM tasks WHERE id = $1;", id)
    if err != nil {
        tx.Rollback()
        return err
    }
    commitErr := tx.Commit()
	return commitErr
}

func (repo *repository) GetTaskById(id int64) (*Task, error) {
    task := NewDefaultTask()
	row := repo.db.QueryRow("SELECT * FROM tasks WHERE id = $1;", id)
    err := row.Scan(&task.Id, &task.Name, &task.Description, &task.Completed, &task.Priority)
    if err == sql.ErrNoRows {
        return nil, ErrorTaskNotFound
    }
    if err != nil {
        return nil, err
    }
    return task, nil
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

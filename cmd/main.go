package main

import (
	"database/sql"
	"log"

    _ "modernc.org/sqlite"

	"github.com/lincolngondin/desafio-junior-backend-simplify/config"
	"github.com/lincolngondin/desafio-junior-backend-simplify/internal/task"
)

func main() {
	config := config.New()
	db, err := sql.Open(config.DBDriverName, config.DBDataSourceName)
	if err != nil {
		log.Fatal(err)
	}
    defer db.Close()
    
    taskRepo := task.NewReposity(db)
    task := task.NewTask("Teste", "", task.TASK_PRIORITY_LOW)
    taskRepo.CreateTask(task)


}

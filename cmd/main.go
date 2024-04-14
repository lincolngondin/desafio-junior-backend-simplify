package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	_ "modernc.org/sqlite"

	"github.com/lincolngondin/desafio-junior-backend-simplify/config"
	"github.com/lincolngondin/desafio-junior-backend-simplify/internal/task"

	"net/http"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	config := config.New()
	db, err := sql.Open(config.DBDriverName, config.DBDataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	taskRepo := task.NewReposity(db)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /task", taskHandler.GetTasksHandler)
	mux.HandleFunc("GET /task/{task_id}", taskHandler.GetTasksHandler)
	mux.HandleFunc("POST /task", taskHandler.CreateTaskHandler)
	mux.HandleFunc("PUT /task/{task_id}", taskHandler.UpdateTaskHandler)
	mux.HandleFunc("DELETE /task/{task_id}", taskHandler.DeleteTaskHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	serverClosed := make(chan bool)

	go func(srv *http.Server, srvClose chan<- bool) {
		log.Println("Listening in ", srv.Addr)
		err := srv.ListenAndServe()

		if err != nil {
			log.Println("Server has been closed!")
		}
		srvClose <- true
	}(&server, serverClosed)

	<-c
	server.Shutdown(context.Background())
	<-serverClosed

}

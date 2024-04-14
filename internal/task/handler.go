package task

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type serv interface {
	GetTasks(id int64, byID bool) ([]Task, error)
	CreateTask(name string, description string, priority Priority) error
	DeleteTask(id int64) error
	UpdateTask(id int64, name, description string, priority Priority) error
}

type handler struct {
	service serv
}

func NewHandler(service serv) *handler {
	return &handler{
		service: service,
	}
}

type ResponseErrors struct {
	Error string `json:"erro"`
}

var ErrFetchingData = ResponseErrors{"Erro interno!"}
var ErrInvalidID = ResponseErrors{"ID inválido!"}
var ErrIDNotFound = ResponseErrors{"ID não encontrado!"}

func (hnd *handler) GetTasksHandler(response http.ResponseWriter, request *http.Request) {
	setId := request.PathValue("task_id")
	encoder := json.NewEncoder(response)
	if setId == "" {
		tasks, err := hnd.service.GetTasks(0, false)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(err)
			return
		}
		response.WriteHeader(http.StatusOK)
		encoder.Encode(tasks)
	} else {
		id, errConv := strconv.ParseInt(setId, 10, 64)
		if errConv != nil {
			response.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ErrInvalidID)
			return
		}
		tasks, err := hnd.service.GetTasks(id, true)
		if err == ErrorTaskNotFound {
			response.WriteHeader(http.StatusNotFound)
			encoder.Encode(ErrIDNotFound)
			return
		}
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(ErrFetchingData)
			return
		}
		response.WriteHeader(http.StatusOK)
		encoder.Encode(tasks[0])
	}
}

type TaskResponse struct {
	Id          int64    `json:"-"`
	Name        string   `json:"nome"`
	Description string   `json:"descricao"`
	Completed   bool     `json:"realizado"`
	Priority    Priority `json:"prioridade"`
}

func (hnd *handler) CreateTaskHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	task := NewDefaultTask()
	decodeErr := decoder.Decode(&task)
	fmt.Println(task)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		response.Write([]byte("Invalid request!"))
	}
	if !task.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid request: nome, descricao e prioridade não podem estar vazios!"))
		return
	}

	err := hnd.service.CreateTask(task.Name, task.Description, task.Priority)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
	} else {
		response.WriteHeader(http.StatusOK)
		encode := json.NewEncoder(response)
		encode.Encode(TaskResponse(*task))
	}
}

func (hnd *handler) UpdateTaskHandler(response http.ResponseWriter, request *http.Request) {
}

func (hnd *handler) DeleteTaskHandler(response http.ResponseWriter, request *http.Request) {
    encoder := json.NewEncoder(response)
    idStr := request.PathValue("task_id")
    id, errConv := strconv.ParseInt(idStr, 10, 64)
    if errConv != nil {
        response.WriteHeader(http.StatusBadRequest)
        encoder.Encode(ErrInvalidID)
        return
    }
    err := hnd.service.DeleteTask(id)
    if err == ErrorTaskNotFound {
        response.WriteHeader(http.StatusNotFound)
        encoder.Encode(ErrIDNotFound)
        return
    }
    if err != nil {
        log.Println(err)
        response.WriteHeader(http.StatusInternalServerError)
        encoder.Encode(ErrFetchingData)
        return
    }
    response.WriteHeader(http.StatusNoContent)
}

package task

type Priority int

const (
	TASK_PRIORITY_NO_PRIORITY = 0
	TASK_PRIORITY_LOW         = 1
	TASK_PRIORITY_MEDIUM      = 2
	TASK_PRIORITY_HIGH        = 3
)

type Task struct {
	Id          int64    `json:"id"`
	Name        string   `json:"nome"`
	Description string   `json:"descricao"`
	Completed   bool     `json:"realizado"`
	Priority    Priority `json:"prioridade"`
}

func NewTask(name string, description string, priority Priority) *Task {
	return &Task{
		Name:        name,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}
}

func NewDefaultTask() *Task {
	return &Task{
		Id:          -1,
		Name:        "",
		Description: "",
		Completed:   false,
		Priority:    TASK_PRIORITY_NO_PRIORITY,
	}
}

func (task *Task) IsValid() bool {
	return task.Name != "" && task.Description != "" && task.Priority != TASK_PRIORITY_NO_PRIORITY
}

func (task *Task) ChangeCompleted() {
	task.Completed = !task.Completed
}

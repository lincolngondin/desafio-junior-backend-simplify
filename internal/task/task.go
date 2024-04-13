package task

type Priority int

const (
	TASK_PRIORITY_LOW    = 0
	TASK_PRIORITY_MEDIUM = 1
	TASK_PRIORITY_HIGH   = 2
)

type Task struct {
	Id          int64
	Name        string
	Description string
	Completed   bool
	Priority    Priority
}

func NewTask(name string, description string, priority Priority) *Task {
	return &Task{
		Name:        name,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}
}

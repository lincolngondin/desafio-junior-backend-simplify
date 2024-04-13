package task

type repo interface {
	CreateTask(task *Task) error
	DeleteTask(id int64) error
	GetTasks() ([]Task, error)
	UpdateTask(id int64, name string, description string, priority Priority, completed bool) error
}

type service struct {
	repository repo
}

func NewService(repository repo) *service {
	return &service{
		repository: repository,
	}
}

func (svc *service) CreateTask() {}

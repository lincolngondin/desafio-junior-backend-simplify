package task

type repo interface {
	CreateTask(task *Task) error
	DeleteTask(id int64) error
    GetTaskById(id int64) (*Task, error)
	GetTasks() ([]Task, error)
	UpdateTask(id int64, name string, description string, priority Priority, completed bool) (*Task, error)
}

type service struct {
	repository repo
}

func NewService(repository repo) *service {
	return &service{
		repository: repository,
	}
}

func (svc *service) GetTasks(id int64, byID bool) ([]Task, error) {
    if byID {
        task, err := svc.repository.GetTaskById(id)
        if err != nil {
            return nil, err
        }
        return []Task{*task}, nil
    } else {
        return svc.repository.GetTasks()
    }
}

func (svc *service) CreateTask(name string, description string, priority Priority) error {
    task := NewTask(name, description, priority)
    return svc.repository.CreateTask(task)
}

func (svc *service) DeleteTask(id int64) error {
    return svc.repository.DeleteTask(id)
}

func (svc *service) UpdateTask(id int64, name, description string, priority Priority, completed bool) (*Task, error) {
    return svc.repository.UpdateTask(id, name, description, priority, completed)
}

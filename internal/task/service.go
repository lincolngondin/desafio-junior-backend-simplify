package task

type repo interface {
    CreateTask(task *Task) error
    DeleteTask(id int64) error
    GetTasks() []Task
    UpdateTask(id int64, name string, description string, priority Priority, completed bool) error
}

type serv interface {
    
}

type service struct {
    repository repo 
}

func NewService(repository repo) *service {
    return &service{
        repository: repository,
    }
}

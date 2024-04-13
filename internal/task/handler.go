package task

type serv interface {
        
}

type handler struct {
    service serv
}

func NewHandler(service serv) *handler {
    return &handler{
        service: service,
    }
}



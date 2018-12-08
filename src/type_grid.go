package src

type server interface {
	Connect() error
	Disconnect() error
	Listen() error
	Notify(callback *NotifyCallback)
	RegisterMethod(method string, handler *RequestHandler) error
	Register(methods *QUID) error
}

// Grid whatever
type Grid struct {
	server
	internalQUID
	InternalQUID
	listenHandlers []*ListenCallback
}

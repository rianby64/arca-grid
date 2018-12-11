package grid

type server interface {
	Listen() error
	Notify(callback *NotifyCallback)
	RegisterMethod(method string, handler *RequestHandler) error
	Register(methods *QUID) error
}

// Grid whatever
type Grid struct {
	server
	internalQUID
	quid
	listenHandlers []*ListenCallback
}

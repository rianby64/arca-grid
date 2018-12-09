package grid

type server interface {
	Listen() error
	notify(callback *NotifyCallback)
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

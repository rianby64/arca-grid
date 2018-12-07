package src

type server interface {
	Connect() error
	Disconnect() error
	Listen() error
	Notify(callback *NotifyCallback)
	RegisterMethod(method string, handler *RequestHandler) error
	Register(methods *InternalQUID) error
}

// Grid whatever
type Grid struct {
	server
	QUID
	InternalQUID
	notifyHandlers []*NotifyCallback
}

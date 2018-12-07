package src

type server interface {
	Connect() error
	Disconnect() error
	Listen() error
	Notify(callback *NotifyCallback)
	RegisterMethod(method string, handler *RequestHandler) error
	Register(methods *InternalSIDU) error
}

// Grid whatever
type Grid struct {
	server
	SIDU
	InternalSIDU
	notifyHandlers []*NotifyCallback
}

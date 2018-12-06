package src

type server interface {
	Connect() error
	Disconnect() error
	Listen() error
	Notify(callback *NotifyCallback)
	Register(handler *RequestHandler) error
}

// Grid whatever
type Grid struct {
	server
	SIDU
	internalSIDU
	notifyHandlers []*NotifyCallback
}

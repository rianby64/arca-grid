package src

// NotifyCallback whatever
type NotifyCallback func(message interface{})

// RequestHandler whatever
type RequestHandler func(requestParams *interface{},
	context *interface{}, notify NotifyCallback) (interface{}, error)

// QUID whatever
type QUID struct {
	Query  *RequestHandler
	Update *RequestHandler
	Insert *RequestHandler
	Delete *RequestHandler
}

type internalQUID struct {
	query  *RequestHandler
	update *RequestHandler
	insert *RequestHandler
	delete *RequestHandler
}

// InternalQUID whatever
type InternalQUID interface {
	Query()
	Update()
	Insert()
	Delete()
}

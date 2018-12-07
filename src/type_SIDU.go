package src

// NotifyCallback whatever
type NotifyCallback func(message interface{})

// RequestHandler whatever
type RequestHandler func(requestParams *interface{},
	context *interface{}, notify NotifyCallback) (interface{}, error)

// InternalQUID whatever
type InternalQUID struct {
	query  *RequestHandler
	update *RequestHandler
	insert *RequestHandler
	delete *RequestHandler
}

// QUID whatever
type QUID interface {
	Query()
	Update()
	Insert()
	Delete()
}

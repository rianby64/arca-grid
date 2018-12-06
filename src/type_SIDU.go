package src

// NotifyCallback whatever
type NotifyCallback func(message interface{})

// RequestHandler whatever
type RequestHandler func(requestParams *interface{},
	context *interface{}, notify NotifyCallback)

type internalSIDU struct {
	selectInternal *RequestHandler
	insertInternal *RequestHandler
	deleteInternal *RequestHandler
	updateInternal *RequestHandler
}

// SIDU whatever
type SIDU interface {
	Select()
	Insert()
	Delete()
	Update()
}

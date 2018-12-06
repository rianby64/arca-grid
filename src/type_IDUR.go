package src

// NotifyCallback whatever
type NotifyCallback func(message interface{})

// internalIDU whatever
type internalIDU struct {
	insertInternal *RequestHandler
	deleteInternal *RequestHandler
	updateInternal *RequestHandler
}

type internalSIDU struct {
	selectInternal *RequestHandler
	internalIDU
}

// RequestHandler whatever
type RequestHandler func(requestParams *interface{},
	context *interface{}, notify NotifyCallback)

// IDU whatever
type IDU interface {
	Insert()
	Delete()
	Update()
}

// SIDU whatever
type SIDU interface {
	Select()
	IDU
}

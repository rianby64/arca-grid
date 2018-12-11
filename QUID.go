package grid

// Query whatever
func (g *Grid) Query(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify ListenCallback = g.Notify
	return callInternal(g.query, &notify, requestParams, context)
}

// Update whatever
func (g *Grid) Update(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify ListenCallback = g.Notify
	return callInternal(g.update, &notify, requestParams, context)
}

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify ListenCallback = g.Notify
	return callInternal(g.insert, &notify, requestParams, context)
}

// Delete whatever
func (g *Grid) Delete(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify ListenCallback = g.Notify
	return callInternal(g.delete, &notify, requestParams, context)
}

func callInternal(
	handler *RequestHandler,
	notify *ListenCallback,
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	done := make(chan bool)
	var Notify NotifyCallback = func(message interface{}) {
		var ctx interface{}
		if context != nil {
			ctx = *context
		}
		(*notify)(message, ctx)
	}
	go (func() {
		result, err = (*handler)(requestParams, context, Notify)
		done <- true
	})()
	<-done
	return
}

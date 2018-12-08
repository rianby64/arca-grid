package src

// Query whatever
func (g *Grid) Query(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify NotifyCallback = g.notify
	return callInternal(g.query, &notify, requestParams, context)
}

// Update whatever
func (g *Grid) Update(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify NotifyCallback = g.notify
	return callInternal(g.update, &notify, requestParams, context)
}

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify NotifyCallback = g.notify
	return callInternal(g.insert, &notify, requestParams, context)
}

// Delete whatever
func (g *Grid) Delete(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	var notify NotifyCallback = g.notify
	return callInternal(g.delete, &notify, requestParams, context)
}

func callInternal(
	handler *RequestHandler,
	notify *NotifyCallback,
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	done := make(chan bool)
	go (func() {
		result, err = (*handler)(requestParams, context, *notify)
		done <- true
	})()
	<-done
	return
}

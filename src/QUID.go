package src

// Query whatever
func (g *Grid) Query(
	requestParams *interface{},
	context *interface{},
) (interface{}, error) {
	go (*g.query)(requestParams, context, g.Notify)
	return nil, nil
}

// Update whatever
func (g *Grid) Update(
	requestParams *interface{},
	context *interface{},
) (result interface{}, err error) {
	done := make(chan bool)
	go (func() {
		result, err = (*g.update)(requestParams, context, g.Notify)
		done <- true
	})()
	<-done
	return
}

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) (interface{}, error) {
	go (*g.insert)(requestParams, context, g.Notify)
	return nil, nil
}

// Delete whatever
func (g *Grid) Delete(
	requestParams *interface{},
	context *interface{},
) (interface{}, error) {
	go (*g.delete)(requestParams, context, g.Notify)
	return nil, nil
}

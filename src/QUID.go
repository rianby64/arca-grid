package src

// Query whatever
func (g *Grid) Query(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.query)(requestParams, context, g.Notify)
	return nil
}

// Update whatever
func (g *Grid) Update(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.update)(requestParams, context, g.Notify)
	return nil
}

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.insert)(requestParams, context, g.Notify)
	return nil
}

// Delete whatever
func (g *Grid) Delete(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.delete)(requestParams, context, g.Notify)
	return nil
}

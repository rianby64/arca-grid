package src

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.insertInternal)(requestParams, context, g.Notify)
	return nil
}

// Delete whatever
func (g *Grid) Delete(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.deleteInternal)(requestParams, context, g.Notify)
	return nil
}

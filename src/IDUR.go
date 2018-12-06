package src

// Insert whatever
func (g *Grid) Insert(
	requestParams *interface{},
	context *interface{},
) error {
	go (*g.insertInternal)(requestParams, context, g.Notify)
	return nil
}

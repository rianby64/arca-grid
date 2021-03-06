package grid

// Listen whatever
func (g *Grid) Listen(callback *ListenCallback) error {
	g.listenHandlers = append(g.listenHandlers, callback)
	return nil
}

// Notify whatever
func (g *Grid) Notify(message interface{}, context interface{}) {
	for _, handler := range g.listenHandlers {
		go (*handler)(message, context)
	}
}

// RegisterMethod whatever
func (g *Grid) RegisterMethod(method string, handler *RequestHandler) error {
	if method == "query" {
		g.query = handler
	}
	if method == "update" {
		g.update = handler
	}
	if method == "insert" {
		g.insert = handler
	}
	if method == "delete" {
		g.delete = handler
	}
	return nil
}

// Register whatever
func (g *Grid) Register(methods *QUID) error {
	g.query = methods.Query
	g.update = methods.Update
	g.insert = methods.Insert
	g.delete = methods.Delete
	return nil
}

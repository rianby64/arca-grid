package src

// Connect whatever
func (g *Grid) Connect(to *Grid) error {
	return nil
}

// Disconnect whatever
func (g *Grid) Disconnect() error {
	return nil
}

// Listen whatever
func (g *Grid) Listen(callback *NotifyCallback) error {
	g.notifyHandlers = append(g.notifyHandlers, callback)
	return nil
}

// Notify whatever
func (g *Grid) Notify(message interface{}) {
	for _, handler := range g.notifyHandlers {
		(*handler)(message)
	}
}

// Register whatever
func (g *Grid) Register(method string, handler *RequestHandler) error {
	g.insertInternal = handler
	return nil
}

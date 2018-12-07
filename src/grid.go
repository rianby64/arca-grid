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
		go (*handler)(message)
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
func (g *Grid) Register(methods *InternalQUID) error {
	g.query = methods.query
	g.update = methods.update
	g.insert = methods.insert
	g.delete = methods.delete
	return nil
}

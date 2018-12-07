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
func (g *Grid) RegisterMethod(method string, handler *RequestHandler) error {
	if method == "select" {
		g.selectInternal = handler
	}
	if method == "insert" {
		g.insertInternal = handler
	}
	if method == "delete" {
		g.deleteInternal = handler
	}
	if method == "update" {
		g.updateInternal = handler
	}
	return nil
}

func (g *Grid) Register(methods *InternalSIDU) error {
	return nil
}

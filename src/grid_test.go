package src

import (
	"testing"
)

func Test_Listen(t *testing.T) {
	t.Log("Test notify from InsertDefinition")

	// Setup
	server := Grid{}
	var listener NotifyCallback = func(message interface{}) {
	}

	// Excercise
	server.Listen(&listener)

	// Verify
	if len(server.notifyHandlers) == 0 {
		t.Errorf("handlers '%v' is empty", server.notifyHandlers)
	}
}

func Test_RegisterMethod_select(t *testing.T) {
	t.Log("Test RegisterMethod select definition")

	// Setup
	var selectDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (*interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("select", &selectDefinition)

	// Verify
	if &selectDefinition != server.selectInternal {
		t.Errorf("selectDefinition '%v' differs from internal select '%v'",
			*server.selectInternal, &selectDefinition)
	}
}

func Test_RegisterMethod_insert(t *testing.T) {
	t.Log("Test RegisterMethod insert definition")

	// Setup
	var insertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (*interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("insert", &insertDefinition)

	// Verify
	if &insertDefinition != server.insertInternal {
		t.Errorf("insertDefinition '%v' differs from internal insert '%v'",
			*server.insertInternal, &insertDefinition)
	}
}

func Test_RegisterMethod_delete(t *testing.T) {
	t.Log("Test RegisterMethod delete definition")

	// Setup
	var deleteDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (*interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("delete", &deleteDefinition)

	// Verify
	if &deleteDefinition != server.deleteInternal {
		t.Errorf("deleteDefinition '%v' differs from internal delete '%v'",
			*server.deleteInternal, &deleteDefinition)
	}
}

func Test_RegisterMethod_update(t *testing.T) {
	t.Log("Test RegisterMethod update definition")

	// Setup
	var updateDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (*interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("update", &updateDefinition)

	// Verify
	if &updateDefinition != server.updateInternal {
		t.Errorf("updateDefinition '%v' differs from internal update '%v'",
			*server.updateInternal, &updateDefinition)
	}
}

package src

import (
	"testing"
)

func Test_Listen(t *testing.T) {
	t.Log("Test Listen method")

	// Setup
	server := Grid{}
	var listener ListenCallback = func(
		message interface{},
		context interface{},
	) {
	}

	// Excercise
	server.Listen(&listener)

	// Verify
	if len(server.listenHandlers) == 0 {
		t.Errorf("handlers '%v' is empty", server.listenHandlers)
	}
}

func Test_notify(t *testing.T) {
	t.Log("Test notify method")

	// Setup
	done := make(chan bool)
	listenerReached := false
	server := Grid{}
	var listener ListenCallback = func(
		message interface{},
		context interface{},
	) {
		listenerReached = true
		done <- true
	}

	// Excercise
	server.Listen(&listener)
	server.notify(nil, nil)
	<-done

	// Verify
	if !listenerReached {
		t.Error("listener is unreachable")
	}

}

func Test_RegisterMethod_query(t *testing.T) {
	t.Log("Test RegisterMethod query definition")

	// Setup
	var queryDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("query", &queryDefinition)

	// Verify
	if &queryDefinition != server.query {
		t.Errorf("queryDefinition '%v' differs from internal query '%v'",
			*server.query, &queryDefinition)
	}
}

func Test_RegisterMethod_insert(t *testing.T) {
	t.Log("Test RegisterMethod insert definition")

	// Setup
	var insertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("insert", &insertDefinition)

	// Verify
	if &insertDefinition != server.insert {
		t.Errorf("insertDefinition '%v' differs from internal insert '%v'",
			*server.insert, &insertDefinition)
	}
}

func Test_RegisterMethod_delete(t *testing.T) {
	t.Log("Test RegisterMethod delete definition")

	// Setup
	var deleteDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("delete", &deleteDefinition)

	// Verify
	if &deleteDefinition != server.delete {
		t.Errorf("deleteDefinition '%v' differs from internal delete '%v'",
			*server.delete, &deleteDefinition)
	}
}

func Test_RegisterMethod_update(t *testing.T) {
	t.Log("Test RegisterMethod update definition")

	// Setup
	var updateDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}

	server := Grid{}

	// Excercise
	server.RegisterMethod("update", &updateDefinition)

	// Verify
	if &updateDefinition != server.update {
		t.Errorf("updateDefinition '%v' differs from internal update '%v'",
			*server.update, &updateDefinition)
	}
}

func Test_Register_methods(t *testing.T) {
	t.Log("Test Register methods")

	// Setup
	var query RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}
	var update RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}
	var insert RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}
	var delete RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {
		return nil, nil
	}
	var methods = &QUID{
		Query:  &query,
		Update: &update,
		Insert: &insert,
		Delete: &delete,
	}

	server := Grid{}

	// Excercise
	server.Register(methods)

	// Verify
	if methods.Query != server.query ||
		methods.Update != server.update ||
		methods.Insert != server.insert ||
		methods.Delete != server.delete {
		t.Error("methods didn't register properly")
	}
}

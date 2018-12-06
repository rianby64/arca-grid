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

func Test_Register_insert(t *testing.T) {
	t.Log("Test register insert definition")

	// Setup
	var insertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) {
	}

	server := Grid{}

	// Excercise
	server.Register("insert", &insertDefinition)

	// Verify
	if &insertDefinition != server.insertInternal {
		t.Errorf("insertDefinition '%v' differs from internal insert '%v'",
			*server.insertInternal, &insertDefinition)
	}
}

func Test_Register_delete(t *testing.T) {
	t.Log("Test register delete definition")

	// Setup
	var deleteDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) {
	}

	server := Grid{}

	// Excercise
	server.Register("delete", &deleteDefinition)

	// Verify
	if &deleteDefinition != server.deleteInternal {
		t.Errorf("deleteDefinition '%v' differs from internal delete '%v'",
			*server.deleteInternal, &deleteDefinition)
	}
}

func Test_Notify_from_insertDefinition(t *testing.T) {
	t.Log("Test notify from insertDefinition")

	// Setup
	ch := make(chan string)
	msgExpected := "message expected"

	var listener NotifyCallback = func(message interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}

		ch <- message.(string)

		// tear down
		close(ch)
	}

	var insertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) {

		// excercise
		notify(msgExpected)
	}

	server := Grid{}
	server.Listen(&listener)
	server.Register("insert", &insertDefinition)

	// Excercise
	server.Insert(nil, nil)

	msgActual, ok := <-ch
	if !ok {
		t.Error("Unexpected error")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

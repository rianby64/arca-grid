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

func Test_Register(t *testing.T) {
	t.Log("Test notify from InsertDefinition")

	// Setup
	var InsertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) {
	}

	server := Grid{}

	// Excercise
	server.Register("", &InsertDefinition)

	// Verify
	if &InsertDefinition != server.insertInternal {
		t.Errorf("InsertDefinition '%v' differs from internal insert '%v'",
			*server.insertInternal, &InsertDefinition)
	}
}

func Test_Notify_from_InsertDefinition(t *testing.T) {
	t.Log("Test notify from InsertDefinition")

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

	var InsertDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) {

		// excercise
		notify(msgExpected)
	}

	server := Grid{}
	server.Listen(&listener)
	server.Register("", &InsertDefinition)

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

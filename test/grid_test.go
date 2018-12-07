package test

import (
	"testing"

	"../src"
)

func Test_Notify_from_queryDefinition(t *testing.T) {
	t.Log("Test notify from queryDefinition")

	// Setup
	ch := make(chan string)
	msgExpected := "message expected"

	var listener src.NotifyCallback = func(message interface{}) {

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

	var queryDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("query", &queryDefinition)

	// Excercise
	server.Query(nil, nil)

	msgActual, ok := <-ch
	if !ok {
		t.Error("Unexpected error")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notify_from_insertDefinition(t *testing.T) {
	t.Log("Test notify from insertDefinition")

	// Setup
	ch := make(chan string)
	msgExpected := "message expected"

	var listener src.NotifyCallback = func(message interface{}) {

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

	var insertDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("insert", &insertDefinition)

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

func Test_Notify_from_deleteDefinition(t *testing.T) {
	t.Log("Test notify from deleteDefinition")

	// Setup
	ch := make(chan string)
	msgExpected := "message expected"

	var listener src.NotifyCallback = func(message interface{}) {

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

	var deleteDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("delete", &deleteDefinition)

	// Excercise
	server.Delete(nil, nil)

	msgActual, ok := <-ch
	if !ok {
		t.Error("Unexpected error")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notify_from_updateDefinition(t *testing.T) {
	t.Log("Test notify from updateDefinition")

	// Setup
	var msgActual string
	msgExpected := "message expected"

	var listener src.NotifyCallback = func(message interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}

		msgActual = message.(string)
	}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	server.Update(nil, nil)
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notifications_from_updateDefinition(t *testing.T) {
	t.Log("Test notifications from updateDefinition")

	// Setup
	var msgActual1 string
	var msgActual2 string
	msgExpected := "message expected"

	var listener1 src.NotifyCallback = func(message interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual1 = message.(string)
	}

	var listener2 src.NotifyCallback = func(message interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual2 = message.(string)
	}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener1)
	server.Listen(&listener2)
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	server.Update(nil, nil)

	if msgActual1 != msgExpected {
		t.Error("received message differs from the expected")
	}

	if msgActual2 != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_result_from_update(t *testing.T) {
	t.Log("Test result from update")

	// Setup
	resultExpected := []string{"a complex result"}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		return resultExpected, nil
	}

	server := src.Grid{}
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	resultActual, err := server.Update(nil, nil)
	if err != nil {
		t.Error("Unexpected error")
	}
	if resultActual == nil {
		t.Error("Action server.Update returned nil")
	}
}

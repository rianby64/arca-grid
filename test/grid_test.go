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

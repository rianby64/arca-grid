package test

import (
	"testing"

	"../src"
)

func Test_Notify_from_selectDefinition(t *testing.T) {
	t.Log("Test notify from selectDefinition")

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

	var selectDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) {

		// excercise
		notify(msgExpected)
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.Register("select", &selectDefinition)

	// Excercise
	server.Select(nil, nil)

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
	) {

		// excercise
		notify(msgExpected)
	}

	server := src.Grid{}
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
	) {

		// excercise
		notify(msgExpected)
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.Register("delete", &deleteDefinition)

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

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) {

		// excercise
		notify(msgExpected)
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.Register("update", &updateDefinition)

	// Excercise
	server.Update(nil, nil)

	msgActual, ok := <-ch
	if !ok {
		t.Error("Unexpected error")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

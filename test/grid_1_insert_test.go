package test

import (
	"testing"

	"../src"
)

func Test_Notify_from_insertDefinition(t *testing.T) {
	t.Log("Test notify from insertDefinition")

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

	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notifications_from_insertDefinition(t *testing.T) {
	t.Log("Test notifications from insertDefinition")

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
	server.Listen(&listener1)
	server.Listen(&listener2)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	server.Insert(nil, nil)

	if msgActual1 != msgExpected {
		t.Error("received message differs from the expected")
	}

	if msgActual2 != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_result_from_insert(t *testing.T) {
	t.Log("Test result from insert")

	// Setup
	msgExpected := "a complex result"

	var insertDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// excercise
		return msgExpected, nil
	}

	server := src.Grid{}
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	msgActual, err := server.Insert(nil, nil)
	if err != nil {
		t.Error("Unexpected error")
	}
	if msgActual == nil {
		t.Error("Action server.insert returned nil")
	}
	if msgActual != msgExpected {
		t.Error("result message differs from the expected")
	}
}

func Test_result_from_insert_and_notify(t *testing.T) {
	t.Log("Test result from insert")

	// Setup
	var msgActual string
	msgExpected := "a complex result"

	var listener src.NotifyCallback = func(message interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual = message.(string)
	}

	var insertDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return msgExpected, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	resultActual, err := server.Insert(nil, nil)
	if err != nil {
		t.Error("Unexpected error")
	}
	if resultActual == nil {
		t.Error("Action server.insert returned nil")
	}
	if resultActual != msgExpected {
		t.Error("result message differs from the expected")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

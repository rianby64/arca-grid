package test

import (
	"testing"

	"../src"
)

func Test_Notify_from_updateDefinition(t *testing.T) {
	t.Log("Test notify from updateDefinition")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "message expected"

	var listener src.ListenCallback = func(
		message interface{}, context interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}

		msgActual = message.(string)
		done <- true
	}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	server.Update(nil, nil)
	<-done

	// Verify
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notifications_from_updateDefinition(t *testing.T) {
	t.Log("Test notifications from updateDefinition")

	// Setup
	done1 := make(chan bool)
	done2 := make(chan bool)
	var msgActual1 string
	var msgActual2 string
	msgExpected := "message expected"

	var listener1 src.ListenCallback = func(
		message interface{}, context interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual1 = message.(string)
		done1 <- true
	}

	var listener2 src.ListenCallback = func(
		message interface{}, context interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual2 = message.(string)
		done2 <- true
	}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return nil, nil
	}

	server := src.Grid{}
	server.Listen(&listener1)
	server.Listen(&listener2)
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	server.Update(nil, nil)
	<-done1
	<-done2

	// Verify
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
	msgExpected := "a complex result"

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		return msgExpected, nil
	}

	server := src.Grid{}
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	msgActual, err := server.Update(nil, nil)

	// Verify
	if err != nil {
		t.Error("Unexpected error")
	}
	if msgActual == nil {
		t.Error("Action server.Update returned nil")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_result_from_update_and_notify(t *testing.T) {
	t.Log("Test result from update")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "a complex result"

	var listener src.ListenCallback = func(
		message interface{}, context interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual = message.(string)
		done <- true
	}

	var updateDefinition src.RequestHandler = func(
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
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	resultActual, err := server.Update(nil, nil)
	<-done

	if err != nil {
		t.Error("Unexpected error")
	}
	if resultActual == nil {
		t.Error("Action server.Update returned nil")
	}
	if resultActual != msgExpected {
		t.Error("result message differs from the expected")
	}
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_result_from_2updates_and_notify(t *testing.T) {
	t.Log("Test result from two updates and notify")

	// Setup
	var msgActual1 interface{}
	var msgActual2 interface{}
	var msgExpected1 interface{} = "a complex result 1"
	var msgExpected2 interface{} = "a complex result 2"

	var listener src.ListenCallback = func(
		message interface{}, context interface{}) {

		// Verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) == msgExpected1.(string) {
			msgActual1 = message
		}
		if message.(string) == msgExpected2.(string) {
			msgActual2 = message
		}
	}

	var updateDefinition src.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify src.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(*requestParams)
		return *requestParams, nil
	}

	server := src.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("update", &updateDefinition)

	// Excercise
	msgActual1, err1 := server.Update(&msgExpected1, nil)
	msgActual2, err2 := server.Update(&msgExpected2, nil)

	if err1 != nil {
		t.Error("Unexpected error")
	}
	if msgActual1 == nil {
		t.Error("Action server.update returned nil")
	}
	if msgActual1.(string) != msgExpected1.(string) {
		t.Error("Result message differs from the expected")
	}
	if msgActual1.(string) != msgExpected1.(string) {
		t.Error("received message differs from the expected")
	}

	if err2 != nil {
		t.Error("Unexpected error")
	}
	if msgActual2 == nil {
		t.Error("Action server.update returned nil")
	}
	if msgActual2.(string) != msgExpected2.(string) {
		t.Error("result message differs from the expected")
	}
	if msgActual2.(string) != msgExpected2.(string) {
		t.Error("received message differs from the expected")
	}
}

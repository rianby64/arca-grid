package grid_test

import (
	"testing"

	grid "github.com/rianby64/arca-grid"
)

func Test_Notify_from_insertDefinition(t *testing.T) {
	t.Log("Test notify from insertDefinition")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "message expected"

	var listener grid.ListenCallback = func(
		message interface{}, context interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}

		msgActual = message.(string)
		done <- true
	}

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := grid.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	server.Insert(nil, nil)
	<-done

	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

func Test_Notifications_from_insertDefinition(t *testing.T) {
	t.Log("Test notifications from insertDefinition")

	// Setup
	done1 := make(chan bool)
	done2 := make(chan bool)
	var msgActual1 string
	var msgActual2 string
	msgExpected := "message expected"

	var listener1 grid.ListenCallback = func(
		message interface{}, context interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual1 = message.(string)
		done1 <- true
	}

	var listener2 grid.ListenCallback = func(
		message interface{}, context interface{}) {

		// verify
		if message == nil {
			t.Error("received message is nil")
		}
		if message.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}
		msgActual2 = message.(string)
		done2 <- true
	}

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// excercise
		notify(msgExpected)
		return nil, nil
	}

	server := grid.Grid{}
	server.Listen(&listener1)
	server.Listen(&listener2)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	server.Insert(nil, nil)
	<-done1
	<-done2

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

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// excercise
		return msgExpected, nil
	}

	server := grid.Grid{}
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
	t.Log("Test result from insert and notify")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "a complex result"

	var listener grid.ListenCallback = func(
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

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return msgExpected, nil
	}

	server := grid.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	resultActual, err := server.Insert(nil, nil)
	<-done

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

func Test_result_from_2inserts_and_notify(t *testing.T) {
	t.Log("Test result from two inserts and notify")

	// Setup
	var msgActual1 interface{}
	var msgActual2 interface{}
	var msgExpected1 interface{} = "a complex result 1"
	var msgExpected2 interface{} = "a complex result 2"

	var listener grid.ListenCallback = func(
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

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(*requestParams)
		return *requestParams, nil
	}

	server := grid.Grid{}
	server.Listen(&listener)
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	msgActual1, err1 := server.Insert(&msgExpected1, nil)
	msgActual2, err2 := server.Insert(&msgExpected2, nil)

	if err1 != nil {
		t.Error("Unexpected error")
	}
	if msgActual1 == nil {
		t.Error("Action server.insert returned nil")
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
		t.Error("Action server.insert returned nil")
	}
	if msgActual2.(string) != msgExpected2.(string) {
		t.Error("result message differs from the expected")
	}
	if msgActual2.(string) != msgExpected2.(string) {
		t.Error("received message differs from the expected")
	}
}

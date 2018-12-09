package test

import (
	"testing"

	"../src"
)

func Test_listen_insert_with_context(t *testing.T) {
	t.Log("Test result from cascade insert with 2 servers")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "message expected"

	var listener src.ListenCallback = func(
		_ interface{}, context interface{}) {

		// Verify
		if context == nil {
			t.Error("received message is nil")
			done <- true
			return
		}
		if context.(string) != msgExpected {
			t.Error("received message differs from the expected")
		}

		msgActual = context.(string)
		done <- true
	}

	var insertDefinition src.RequestHandler = func(
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
	server.RegisterMethod("insert", &insertDefinition)

	// Excercise
	var context interface{} = msgExpected
	server.Insert(nil, &context)
	<-done

	// Verify
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}

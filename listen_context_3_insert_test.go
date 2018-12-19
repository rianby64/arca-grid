package grid_test

import (
	"testing"

	grid "github.com/rianby64/arca-grid"
)

func Test_listen_insert_with_context(t *testing.T) {
	t.Log("Test listen to insert with context")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "message expected"

	var listener grid.ListenCallback = func(
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

	var insertDefinition grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return nil, nil
	}

	server := grid.Grid{}
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

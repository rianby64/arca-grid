package grid

import (
	"testing"
)

func Test_listen_query_with_context(t *testing.T) {
	t.Log("Test listen to query with context")

	// Setup
	done := make(chan bool)
	var msgActual string
	msgExpected := "message expected"

	var listener ListenCallback = func(
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

	var queryDefinition RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify NotifyCallback,
	) (interface{}, error) {

		// Excercise
		notify(msgExpected)
		return nil, nil
	}

	server := Grid{}
	server.Listen(&listener)
	server.RegisterMethod("query", &queryDefinition)

	// Excercise
	var context interface{} = msgExpected
	server.Query(nil, &context)
	<-done

	// Verify
	if msgActual != msgExpected {
		t.Error("received message differs from the expected")
	}
}
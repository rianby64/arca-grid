package grid_test

import (
	"fmt"
	"testing"

	grid "github.com/rianby64/arca-grid"
)

func Test_result_cascade_query_2_servers(t *testing.T) {
	t.Log("Test result from cascade query with 2 servers")

	server1 := grid.Grid{}
	server2 := grid.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"
	msgExpected := "a custom query + another data"

	var queryHandler1 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Query(requestParams, nil)
		result := iResult.(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		return result, nil
	}

	var queryHandler2 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		return result, nil
	}

	server1.RegisterMethod("query", &queryHandler1)
	server2.RegisterMethod("query", &queryHandler2)

	var iRequest interface{} = request
	result, _ := server1.Query(&iRequest, nil)

	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_query_2_servers(t *testing.T) {
	t.Log("Test notify from cascade query with 2 servers")

	done1 := make(chan bool)
	done2 := make(chan bool)
	server1 := grid.Grid{}
	server2 := grid.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"
	msgExpected := "a custom query + another data"

	var queryHandler1 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Query(requestParams, nil)
		newresult := make(map[string]string)
		result := iResult.(map[string]string)
		for key, value := range result {
			newresult[key] = value
		}
		newresult["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		notify(newresult)
		return newresult, nil
	}

	var listener1 grid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected {
			t.Error("listener1 got a wrong message")
		}
		done1 <- true
	}

	var queryHandler2 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var listener2 grid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != request["query"] {
			t.Error("listener2 got a wrong message")
		}
		done2 <- true
	}

	server1.RegisterMethod("query", &queryHandler1)
	server2.RegisterMethod("query", &queryHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	result, _ := server1.Query(&iRequest, nil)

	<-done1
	<-done2
	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_query_notifier_2_servers(t *testing.T) {
	t.Log("Test notify from cascade query that triggers a notification with 2 servers")

	server1 := grid.Grid{}
	server2 := grid.Grid{}
	done1 := make(chan bool)
	done2 := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"
	msgExpected1 := "a custom query + another data"
	msgExpected2 := "a custom query"

	var queryHandler1 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		request := (*requestParams).(map[string]string)
		result := make(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", request["query"], "another data")
		notify(result)
		return result, nil
	}

	var listener1 grid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected1 {
			t.Error("msgExpected1 differs from the actual")
		}
		done1 <- true
	}

	var queryHandler2 grid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify grid.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var listener2 grid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected2 {
			t.Error("msgExpected2 differs from the actual")
		}

		request := make(map[string]string)
		request["action"] = "do-a-query"
		request["query"] = result["data"]
		var iRequest interface{} = request
		server1.Query(&iRequest, nil)
		done2 <- true
	}

	server1.RegisterMethod("query", &queryHandler1)
	server2.RegisterMethod("query", &queryHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	server2.Query(&iRequest, nil)

	<-done1
	<-done2
}

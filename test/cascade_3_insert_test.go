package test

import (
	"fmt"
	"testing"

	"../src"
)

func Test_result_cascade_insert_2_servers(t *testing.T) {
	t.Log("Test result from cascade insert with 2 servers")

	server1 := src.Grid{}
	server2 := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-insert"
	request["insert"] = "a custom insert"
	msgExpected := "a custom insert + another data"

	var insertHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Insert(requestParams, nil)
		result := iResult.(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		return result, nil
	}

	var insertHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["insert"]
		return result, nil
	}

	server1.RegisterMethod("insert", &insertHandler1)
	server2.RegisterMethod("insert", &insertHandler2)

	var iRequest interface{} = request
	result, _ := server1.Insert(&iRequest, nil)

	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_insert_2_servers(t *testing.T) {
	t.Log("Test notify from cascade insert with 2 servers")

	done1 := make(chan bool)
	done2 := make(chan bool)
	server1 := src.Grid{}
	server2 := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-insert"
	request["insert"] = "a custom insert"
	msgExpected := "a custom insert + another data"

	var insertHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Insert(requestParams, nil)
		newresult := make(map[string]string)
		result := iResult.(map[string]string)
		for key, value := range result {
			newresult[key] = value
		}
		newresult["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		notify(newresult)
		return newresult, nil
	}

	var listener1 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected {
			t.Error("listener1 got a wrong message")
		}
		done1 <- true
	}

	var insertHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["insert"]
		notify(result)
		return result, nil
	}

	var listener2 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		if result["data"] != request["insert"] {
			t.Error("listener2 got a wrong message")
		}
		done2 <- true
	}

	server1.RegisterMethod("insert", &insertHandler1)
	server2.RegisterMethod("insert", &insertHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	result, _ := server1.Insert(&iRequest, nil)

	<-done1
	<-done2
	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_insert_notifier_2_servers(t *testing.T) {
	t.Log("Test notify from cascade insert that triggers a notification with 2 servers")

	server1 := src.Grid{}
	server2 := src.Grid{}
	done1 := make(chan bool)
	done2 := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-a-insert"
	request["insert"] = "a custom insert"
	msgExpected1 := "a custom insert + another data"
	msgExpected2 := "a custom insert"

	var insertHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		request := (*requestParams).(map[string]string)
		result := make(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", request["insert"], "another data")
		notify(result)
		return result, nil
	}

	var listener1 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected1 {
			t.Error("msgExpected1 differs from the actual")
		}
		done1 <- true
	}

	var insertHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["insert"]
		notify(result)
		return result, nil
	}

	var listener2 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected2 {
			t.Error("msgExpected2 differs from the actual")
		}

		request := make(map[string]string)
		request["action"] = "do-a-insert"
		request["insert"] = result["data"]
		var iRequest interface{} = request
		server1.Insert(&iRequest, nil)
		done2 <- true
	}

	server1.RegisterMethod("insert", &insertHandler1)
	server2.RegisterMethod("insert", &insertHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	server2.Insert(&iRequest, nil)

	<-done1
	<-done2
}

package test

import (
	"fmt"
	"testing"

	"../src"
)

func Test_result_cascade_delete_2_servers(t *testing.T) {
	t.Log("Test result from cascade delete with 2 servers")

	server1 := src.Grid{}
	server2 := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-delete"
	request["delete"] = "a custom delete"
	msgExpected := "a custom delete + another data"

	var deleteHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Delete(requestParams, nil)
		result := iResult.(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		return result, nil
	}

	var deleteHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["delete"]
		return result, nil
	}

	server1.RegisterMethod("delete", &deleteHandler1)
	server2.RegisterMethod("delete", &deleteHandler2)

	var iRequest interface{} = request
	result, _ := server1.Delete(&iRequest, nil)

	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_delete_2_servers(t *testing.T) {
	t.Log("Test notify from cascade delete with 2 servers")

	done1 := make(chan bool)
	done2 := make(chan bool)
	server1 := src.Grid{}
	server2 := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-delete"
	request["delete"] = "a custom delete"
	msgExpected := "a custom delete + another data"

	var deleteHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		iResult, _ := server2.Delete(requestParams, nil)
		newresult := make(map[string]string)
		result := iResult.(map[string]string)
		for key, value := range result {
			newresult[key] = value
		}
		newresult["data"] = fmt.Sprintf("%s + %s", result["data"], "another data")
		notify(newresult)
		return newresult, nil
	}

	var listener1 src.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected {
			t.Error("listener1 got a wrong message")
		}
		done1 <- true
	}

	var deleteHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["delete"]
		notify(result)
		return result, nil
	}

	var listener2 src.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != request["delete"] {
			t.Error("listener2 got a wrong message")
		}
		done2 <- true
	}

	server1.RegisterMethod("delete", &deleteHandler1)
	server2.RegisterMethod("delete", &deleteHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	result, _ := server1.Delete(&iRequest, nil)

	<-done1
	<-done2
	if result.(map[string]string)["data"] != msgExpected {
		t.Error("Result data differs from the expected")
	}
}

func Test_notify_from_cascade_delete_notifier_2_servers(t *testing.T) {
	t.Log("Test notify from cascade delete that triggers a notification with 2 servers")

	server1 := src.Grid{}
	server2 := src.Grid{}
	done1 := make(chan bool)
	done2 := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-a-delete"
	request["delete"] = "a custom delete"
	msgExpected1 := "a custom delete + another data"
	msgExpected2 := "a custom delete"

	var deleteHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		request := (*requestParams).(map[string]string)
		result := make(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", request["delete"], "another data")
		notify(result)
		return result, nil
	}

	var listener1 src.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected1 {
			t.Error("msgExpected1 differs from the actual")
		}
		done1 <- true
	}

	var deleteHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["delete"]
		notify(result)
		return result, nil
	}

	var listener2 src.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		if result["data"] != msgExpected2 {
			t.Error("msgExpected2 differs from the actual")
		}

		request := make(map[string]string)
		request["action"] = "do-a-delete"
		request["delete"] = result["data"]
		var iRequest interface{} = request
		server1.Delete(&iRequest, nil)
		done2 <- true
	}

	server1.RegisterMethod("delete", &deleteHandler1)
	server2.RegisterMethod("delete", &deleteHandler2)
	server1.Listen(&listener1)
	server2.Listen(&listener2)
	var iRequest interface{} = request
	server2.Delete(&iRequest, nil)

	<-done1
	<-done2
}

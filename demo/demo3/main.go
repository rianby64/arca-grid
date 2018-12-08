package main

import (
	"fmt"
	"log"

	"../../src"
)

func main() {
	server1 := src.Grid{}
	server2 := src.Grid{}
	done1 := make(chan bool)
	done2 := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"

	var queryHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		var iRequest interface{} = request
		result := iRequest.(map[string]string)
		result["data"] = fmt.Sprintf("%s + %s", result["query"], "another data")
		notify(result)
		return result, nil
	}

	var listener1 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		log.Println(result, "listener1")
		done1 <- true
	}

	var queryHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var listener2 src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		log.Println(result, "listener2")

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

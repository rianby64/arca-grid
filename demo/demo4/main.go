package main

import (
	"log"

	"../../src"
)

func main() {
	server := src.Grid{}
	done := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-something"
	request["query"] = "a custom query"

	var listener src.NotifyCallback = func(message interface{}) {
		result := message.(map[string]string)
		log.Println(result, "notified")
		done <- true
	}

	var queryHandler src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["query"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var updateHandler src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["update"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var insertHandler src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["insert"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var deleteHandler src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["delete"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	methods := src.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	server.Register(&methods)
	server.Listen(&listener)
	var iRequest interface{} = request
	go server.Query(&iRequest, nil)
	go server.Update(&iRequest, nil)
	go server.Insert(&iRequest, nil)
	go server.Delete(&iRequest, nil)

	times := 0
	for {
		<-done
		times++
		if times == 4 {
			break
		}
	}
}

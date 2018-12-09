package main

import (
	"log"

	arcagrid "../.."
)

func main() {
	server := arcagrid.Grid{}
	done := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-something"
	request["query"] = "a custom query"

	var listener arcagrid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		log.Println(result, "notified with context", context)
		done <- true
	}

	var queryHandler arcagrid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify arcagrid.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["query"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var updateHandler arcagrid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify arcagrid.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["update"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var insertHandler arcagrid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify arcagrid.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["insert"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	var deleteHandler arcagrid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify arcagrid.NotifyCallback) (interface{}, error) {
		result := make(map[string]string)
		result["delete"] = (*requestParams).(map[string]string)["query"]
		notify(result)
		return result, nil
	}

	methods := arcagrid.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	server.Register(&methods)
	server.Listen(&listener)
	var iRequest interface{} = request
	var iCtxQuery interface{} = map[string]string{"ctx": "Query"}
	go server.Query(&iRequest, &iCtxQuery)
	var iCtxUpdate interface{} = map[string]string{"ctx": "Update"}
	go server.Update(&iRequest, &iCtxUpdate)
	var iCtxInsert interface{} = map[string]string{"ctx": "Insert"}
	go server.Insert(&iRequest, &iCtxInsert)
	var iCtxDelete interface{} = map[string]string{"ctx": "Delete"}
	go server.Delete(&iRequest, &iCtxDelete)

	times := 0
	for {
		<-done
		times++
		if times == 4 {
			break
		}
	}
}

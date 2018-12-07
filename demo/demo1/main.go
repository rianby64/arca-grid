package main

import (
	"log"

	"../../src"
)

func main() {
	server := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"

	var insertHandler src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		return result, nil
	}

	server.RegisterMethod("query", &insertHandler)
	var request2 interface{} = request
	result, _ := server.Query(&request2, nil)

	log.Println(result)
}

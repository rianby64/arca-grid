package main

import (
	"log"

	arcagrid "../.."
)

func main() {
	server := arcagrid.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"

	var queryHandler arcagrid.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify arcagrid.NotifyCallback) (interface{}, error) {

		result := make(map[string]string)
		result["data"] = (*requestParams).(map[string]string)["query"]
		return result, nil
	}

	server.RegisterMethod("query", &queryHandler)
	var iRequest interface{} = request
	result, _ := server.Query(&iRequest, nil)

	log.Println(result)
}

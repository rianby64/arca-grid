package test

import (
	"fmt"
	"testing"

	"../src"
)

func Test_result_cascade_query_2_servers(t *testing.T) {
	t.Log("Test result from cascade query with 2 servers")

	server1 := src.Grid{}
	server2 := src.Grid{}
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"
	msgExpected := "a custom query + another data"

	var queryHandler1 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

		var iRequest interface{} = request
		iResult2, _ := server2.Query(&iRequest, nil)
		result2 := iResult2.(map[string]string)
		result2["data"] = fmt.Sprintf("%s + %s", result2["data"], "another data")
		return result2, nil
	}

	var queryHandler2 src.RequestHandler = func(requestParams *interface{},
		context *interface{}, notify src.NotifyCallback) (interface{}, error) {

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

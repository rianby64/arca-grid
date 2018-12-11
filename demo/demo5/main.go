package main

import (
	"log"
	"time"

	arcagrid "../.."
)

func main() {
	server1 := arcagrid.Grid{}
	done1 := make(chan bool)
	request := make(map[string]string)
	request["action"] = "do-a-query"
	request["query"] = "a custom query"

	time.AfterFunc(time.Millisecond*100, func() {
		result := map[string]string{
			"data": "some cool data",
		}

		context := map[string]string{
			"source": "a source",
		}

		var iRequest interface{} = result
		var iContext interface{} = context
		server1.Notify(iRequest, iContext)
	})

	var listener1 arcagrid.ListenCallback = func(
		message interface{}, context interface{}) {
		result := message.(map[string]string)
		log.Println(result, "listener1")
		done1 <- true
	}

	server1.Listen(&listener1)

	<-done1
}

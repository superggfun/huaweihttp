package main

import (
	"encoding/json"
	"fmt"
	"main/myProxy"

	"huaweicloud.com/go-runtime/events/apig"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
)

func ApigTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var apigEvent apig.APIGTriggerEvent
	err := json.Unmarshal(payload, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	proxy := myProxy.NewAPIG(apigEvent)
	return proxy.Start()
}

func main() {
	runtime.Register(ApigTest)
}

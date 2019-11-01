package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Code int
	Message string
	Data interface{}
}

func main()  {
	message := Message{200,"success","ok"}
	var result []byte = encodeJson(message)
	decodeJson(result,&message)
	fmt.Println(message.Data)
}

/**
 * json转化
 */
func encodeJson(message Message) []byte  {
	result,err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(result))

	return result
}

/**
 * 解析json
 */
func decodeJson(jsonString []byte,message *Message)  {
	err := json.Unmarshal(jsonString,message)
	if err != nil {
		fmt.Println(err)
	}
}
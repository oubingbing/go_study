package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "hello world"

	result := strings.Index(str,"or")

	fmt.Println(result)
}

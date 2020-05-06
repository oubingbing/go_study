package main

import (
	"fmt"
	"strings"
)

func main()  {
	var result []byte
	reader := strings.NewReader("hello999world")
	c,err := reader.ReadAt(result,0)

	fmt.Println(string(result[:c]),err)
}

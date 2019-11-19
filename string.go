package main

import (
	"fmt"
	"strings"
)

func main()  {
	CompareStringByBuilder()
}

func AppendStringByBuilder()  {
	sb := strings.Builder{}
	sb.Write([]byte("中华人民"))
	sb.WriteString("共和国")
	fmt.Println(sb.String())
}

func CompareStringByBuilder()  {
	str := "Hello"
	result := strings.Compare(str,"Hello")
	fmt.Println(result)
}

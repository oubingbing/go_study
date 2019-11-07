package main

import (
	"fmt"
	"strings"
)

//测试字符串拼接
func main() {
	str := []string{"a","b","c"}

	var newStr strings.Builder
	for _,s := range  str{
		fmt.Fprint(&newStr,s)
	}

	fmt.Println(newStr.String())

	AppendString()
}

//高效的字符串拼接
func AppendString()  {
	hi := "hello "
	w := "world"

	var b strings.Builder
	fmt.Fprint(&b,hi)
	fmt.Fprint(&b,w)

	fmt.Println(b.String())
}
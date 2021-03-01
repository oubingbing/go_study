package main

import (
	"fmt"
)

func main()  {
	var a interface{}
	var c float64
	c = 12
	a = c
	v,ok := a.(float64)
	fmt.Println(v,ok)
}
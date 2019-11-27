package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
	str := strconv.Itoa(12)
	fmt.Println(reflect.TypeOf(str))
}

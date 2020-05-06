package main

import (
	"fmt"
	"reflect"
)

func main()  {
	/*data := 130
	v := reflect.ValueOf(data) // a reflect.Value
	x := v.Interface()      // an interface{}
	fmt.Println(x)*/
	//i := x.(int)            // an int
	//fmt.Printf("%d\n", i)   // "3"


	i := 130

	v := reflect.ValueOf(i)

	fmt.Println(v.Kind())
}

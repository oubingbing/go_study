package main

import (
	"fmt"
	"reflect"
)

type Rest struct {
	A string
	B string
}

type Test struct {
	R *Rest
	T string
}

func main()  {
	t1 := Test{&Rest{"a","b"},"test"}

	t := reflect.ValueOf(&t1)

	rest := t.Elem().Field(0)

	rest.Elem().Field(0).SetString("sdfsd")

	//一下等同于 set方法
	/*p := a.Interface().(*string)
   	*p = "change value"*/

	fmt.Println(t1.R)
}

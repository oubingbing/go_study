package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
}

type person2 struct {
	name string
	age int
}

func main()  {
	var men *Human = new(Human)
	fmt.Println(reflect.TypeOf(men))
	AddAge(men)
	fmt.Println(men.age)
	fmt.Println(&men)

	binary.Read()
}

func AddAge(men *Human)  {
	men.age = 28
}

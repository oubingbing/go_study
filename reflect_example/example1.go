package main

import (
	"fmt"
	"reflect"
)

func main()  {

	setValue()

	
}

func setValue()  {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	theValue := pointer.Elem()//获取到反射对象本身，即num
	fmt.Println(theValue.CanSet())

	theValue.SetFloat(77)

	fmt.Println(num)

}

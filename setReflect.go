package main

import (
	"fmt"
	"reflect"
)

//测试将请求参数绑定到指定类型成员中
func main()  {
	r3()
}

type Ps struct {
	Name string
	Age string
}

func r3()  {
	person := Ps{}

	pv := reflect.ValueOf(&person).Elem()
	//pt := reflect.TypeOf(&person).Elem()
	for i:=0;i<pv.NumField();i++{
		pf := pv.Field(i)
		if !pf.CanSet(){
			fmt.Println("不可寻址")
			continue
		}
		//绑定数据
		pf.SetString("测试")
	}

	fmt.Println(person)
}

func ChangeValue(person *Ps)  {
	pv := reflect.ValueOf(person).Elem()
	pt := reflect.TypeOf(person).Elem()
	for i:=0; i<pt.NumField();i++  {
		f := pv.Field(i)
		if !f.CanSet() {
			fmt.Println("不可寻址")
			continue
		}
		f.SetString("叶子")
	}
}

func r2()  {
	x := 2
	v := reflect.ValueOf(&x).Elem()
	v.Set(reflect.ValueOf(123))
	fmt.Println(x)
}

func r1()  {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)
}

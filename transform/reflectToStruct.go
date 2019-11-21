package main

import (
	"fmt"
	"reflect"
)

func main()  {
	data := map[string]interface{}{"name":"冰冰","age":18,"city":"深圳"}
	var man Man
	mapStruct(&man,data)
}

type Man struct {
	Name string `sql:"name"`
	Age int `sql:"age"`
	City string `sql:"city"`
}

//将数据库数据映射到对应的具体的类型上
func mapStruct(man *Man,data map[string]interface{})  {
	//获取标签
	//设置值
	v := reflect.ValueOf(man)
	t := reflect.TypeOf(man)
	for i:=0; i<v.Elem().NumField();i++  {
		//Elem获取到对应的指针
		tag := t.Elem().Field(i).Tag.Get("sql")
		v.Elem().Field(i).Set(reflect.ValueOf(data[tag]))
	}
	fmt.Println(man.Name)
}
package main

import "fmt"

func main()  {
	a := CalculateType(add)//将add函数强制转换成CalculateType函数类型
	b := CalculateType(mul)//将mul函数强制转化成CalculateType函数类型
	a(1,2)
	b(1,2)
	a.Serve()
	b.Serve()
}

type CalculateType func(int,int) //声明一个函数类型

func (c *CalculateType)Serve()  {
	fmt.Println("我是函数类型实现的函数")
}

func add(a,b int)  {
	fmt.Println(a+b)
}

func mul(a,b int)  {
	fmt.Println(a*b)
}

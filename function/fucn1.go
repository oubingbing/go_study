package main

import "fmt"

type F1 func(int) bool

func f2(i int) bool {
	return i == 1
}

func main()  {
	a := 2
	beforPrint(a, func(i int) bool {
		fmt.Printf("回调函数%v",i)
		return  false
	})
}

func beforPrint(a int,f F1)  {
	fmt.Println("befor")
	defer fmt.Println(f(a))
}

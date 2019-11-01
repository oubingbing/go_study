package main

import "fmt"

func main()  {
	go test1();
	test2();
	//返回结果是test1和test2无顺序打印数据，两个方法同步执行
}

func test1()  {
	for i:=1;i<1000;i++ {
		fmt.Printf("i %v \n",i)
	}
}

func test2()  {
	for j:=100;j<1000;j++{
		fmt.Printf("j %v \n",j)
	}
}

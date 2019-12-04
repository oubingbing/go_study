package main

import "fmt"

func main()  {
	//切片作为可变参数传递
	sumSlice := []int{1,2,3,4,33,22,13}
	sum(sumSlice...)

	sum(2,232,34,445,44)
}

//可变参函数
func sum(i...int)  {
	for _,v := range i {
		fmt.Println(v)
	}
}

func add(str int,j...int)  {

}

func sub(i,j int)  {

}

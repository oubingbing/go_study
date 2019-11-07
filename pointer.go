package main

import "fmt"

func main()  {
/*	var arr [5]int = [5]int{1,2,3,4,5};
	p1,p2 := &arr,&arr[0];
	fmt.Println(p1);//&[1,2,3,4,5] 整个数组的内存地址
	fmt.Println(p2);//0xc0000180c0 数组第一个元素的内存地址
	fmt.Printf("%T",p1);//*[5]int    数组指针
	fmt.Printf("%T",p2);//*int        指针*/

	sli := []int{1,2,3}

	fmt.Println(&sli[0])
}

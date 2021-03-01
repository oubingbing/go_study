package main

import "fmt"

func main() {
	//TestMake()
	/*s := make([]int,3)
	s[0] = 0
	s[1] = 1
	TestA(s)
	fmt.Println(s)*/

	/*mp := make(map[int]int)
	mp[0] = 0
	mp[1] = 1
	TestB(mp)
	fmt.Println(mp)*/

	arr := [3]int{1,2,3}
	TestC(&arr)
	fmt.Println(arr)
}

func TestC(arr *[3]int)  {
	arr[0] = 111
}

func TestB(mp map[int]int)  {
	mp[0] = 11
}

func TestA(s []int)  {
	s[0] = 11
}

func TestMake()  {
	mp := make(map[string]string)
	mp["a"] = "a"
	fmt.Println(mp)
	fmt.Println(mp)
}

func TestNew()  {
	a := new(int)  //分配了一个内存，并把这个内存的地址赋值给了变量a，值为零值，他就是一个值的指针地址，
	var b *int //声明了一个指针变量b，但是他未指向任何内存地址
	b = a //将new(int)分配的内存地址赋值给b，此时的b指向了new(int)
	*b = 12 //将12写入new(int)出的内存中
	fmt.Println(*a) //打印12
	fmt.Println(*b) //打印12
	fmt.Println(a,b)//打印的将是new(int)分配的那个内存地址

	//所以总的来说，new是分配了一个零值的内存空间
}
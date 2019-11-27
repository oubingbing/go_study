package main

import (
	"fmt"
	"sort"
)

/**
 * 一个类型只要实现了Interface即可使用sort包进行操作
 */
func main()  {
	var data IntByte
	data = []int{3,2,43,5,6}

	//sort.Sort(data) //顺序
	sort.Sort(sort.Reverse(data)) //倒叙

	fmt.Println(data)
}

type IntByte []int

func (i IntByte) Len() int {
	return  len(i)
}

func (ib IntByte) Less(i,j int) bool {
	return ib[i] < ib[j]
}

func (ib IntByte) Swap(i,j int)  {
	ib[j],ib[i] = ib[i],ib[j]
}
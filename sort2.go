package main

import (
	"fmt"
	"sort"
)

func main()  {
	data := SortList{
		{Name:"叶子",Age:20,Birth:"1994"},
		{Name:"宝连",Age:18,Birth:"2000"},
		{Name:"小明",Age:14,Birth:"1990"},
		{Name:"笑话",Age:30,Birth:"1999"},
	}

	sort.Sort(sort.Reverse(SortList(data)))

	fmt.Println(data)
	for _,d := range data{
		fmt.Println(*d)
	}
}

type StudentInfo struct {
	Name string
	Age int
	Birth string
}

type SortList []*StudentInfo

func (s SortList) Len() int {
	return len(s)
}

func (s SortList) Less(i,j int) bool {
	return s[i].Age < s[j].Age
}

func (s SortList) Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}


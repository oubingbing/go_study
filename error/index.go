package main

import "fmt"

func main()  {

	sli := make([]int,1,5)
	sli2 := sli

	t(sli)
	t2(sli2)

	fmt.Println(sli)
	fmt.Println(sli2)

}

func m(mp map[int]int)  {
	mp[1] = 2
}

func t(sli []int) {
	sli[0] = 11
}

func t2(sli []int) {
	sli[0] = 33
}

func t1(sli []int)  {


	//fmt.Println(sli)
	//fmt.Println(cap(sli))
}

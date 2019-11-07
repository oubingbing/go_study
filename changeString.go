package main

import "fmt"

func main()  {
	var str string = "hello";
	//str = "world";
	var sp *string
	sp = &str
	fmt.Println(sp)
}

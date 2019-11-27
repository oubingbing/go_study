package main

import "fmt"

func main()  {
	a,b := 2,3
	fmt.Println(Calculate(a,b,add1))
	fmt.Println(Calculate(a,b,mul1))
}

type CalculateType2 func(a,b int) int

func add1(a,b int) int  {
	return a+b
}

func mul1(a,b int) int {
	return a*b
}

func Calculate(a,b int,f CalculateType2) int {
	return f(a,b)
}

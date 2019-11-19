package test

import (
	"fmt"
	"testing"
)

func main()  {

}

func TestTime(t *testing.T)  {
	var sum int
	for i:=0;i<=40000;i++{
		for j:=0;j<=40000 ;j++  {
			sum += i*j
		}
	}

	fmt.Println(sum)
}

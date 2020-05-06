package main

import "fmt"

type User struct {
	Name string
	Age uint
}

func (this User) String() string  {
	return fmt.Sprintf("my name is %v and age %v",this.Name,this.Age)
}



func main()  {
	user := User{"bingbing",18}
	fmt.Println(user)
}

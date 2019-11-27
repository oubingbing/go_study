package main

func main()  {
	//var student Student
	//student.Write([]byte("bingbing"))

	//student.Write([]byte("ou"))
	/*//fmt.Println(student)

	_,err := fmt.Fprintf(&student,"%s","叶子")
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(student)*/

	var h Human
	h.run()

}

type Human interface {
	run()
	say()
}

type Student struct {
	Name string
	Age int
}

func (s *Student) Write (d []byte) (int,error) {
	s.Name += string(d)
	return len(d),nil
}

package main


import(
	"fmt"
)


type Speaker interface{

	speak()
}

type Human struct{
	Name string
}

func (h Human) speak() {

	fmt.Println("Hello!, My name is ", h.Name)
}


func main() {

	var intro Speaker

	intro = Human{Name: "samuel"}

	intro.speak()

	fmt.Println(intro)
}
package main

import(
	"fmt"
)
//PointOne takes a pointer to an int and change the value to 1
func PointOne(n *int) {

	*n = 1
}
//UltimatePointOne is a pointer to pointer
func UltimatePointOne(n ***int) {
	***n = 1
}

// DivMod divide a and b 
// the result of the division will be stored in the int pointed by div and the remainder will be stored in the int pointed by mod
func DivMod(a int, b int, div *int, mod *int) {
	
	*div = a/b
	*mod = a%b
}

func main() {

	//PointOne 
	n := 0
	PointOne(&n)
	fmt.Println(n)


	a := 0
	b :=&a
	c := &b

	UltimatePointOne(&c)
	fmt.Println(**c)


}

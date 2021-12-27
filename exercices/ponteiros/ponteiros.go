package ponteiros

import "fmt"

func Ponteiros() {
	x := 11

	estarecebeovalor(x)
	estarecebeumponteiro(&x)

	fmt.Println(x)
}
func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função:", x)

}

func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}

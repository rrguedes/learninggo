package main

import "fmt"

func main() {
	x := []string{"Rafael", "Cibelle", "Gustavo", "Gabriel", "Lucas"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

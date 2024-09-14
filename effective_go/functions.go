package main

import (
	"fmt"
)

func nextInt(byteValue []byte, i int) (int, int) {
	for ; i < len(byteValue) && !IsDigit(byteValue[i]); i++ {
	}
	x := 0
	for ; i < len(byteValue) && IsDigit(byteValue[i]); i++ {
		x = x*10 + int(byteValue[i]) - '0'
	}
	return x, i
}

func IsDigit(value byte) bool {
	if value < 48 || value > 57 {
		return false
	}
	return true
}

func main() {
	hello := "Hola Go!"
	var idCliente []byte = []byte(hello)
	x := 0
	for i := 0; i < len(idCliente); {
		x, i = nextInt(idCliente, i)
		fmt.Println(x)
	}
}

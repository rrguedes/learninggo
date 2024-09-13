package main

import (
	"fmt"
)

func main() {
	// Exercicio 1
	greetings := [...]string{"Hello", "Hola", "สวัสดี", "こんにちは", "привет"}
	firstTwoGreetings := greetings[:2]
	intermediareGreetings := greetings[1:3]
	finalGreetings := greetings[3:5]
	fmt.Println("%s : %s", "Slice original", greetings)
	fmt.Println("%s : %s", "Primeiro e Segundo", firstTwoGreetings)
	fmt.Println("%s : %s", "Segundo, Terceiro e Quarto", intermediareGreetings)
	fmt.Println("%s : %s", "Quarto e Quinto", finalGreetings)

	// Exercicio 2
	message := "Hi 🎉 and 🚀"
	fmt.Println("4th element: ", string(message[3]))
	// Exercício 3
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	rafael := Employee{
		"Rafael",
		"Guedes",
		1,
	}

	cibelle := Employee{
		firstName: "Cibelle",
		lastName:  "Guedes",
		id:        2,
	}

	var edson = Employee{
		"Edson",
		"Candido",
		3,
	}
	fmt.Println(rafael)
	fmt.Println(cibelle)
	fmt.Println(edson)
}

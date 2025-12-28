package main

import "fmt"

func inferenceType() {
	fmt.Println("\n -------------------")
	fmt.Println("Variables with inference type")
	name := "Arifian Saputra"
	age := 24
	salary := 150000
	isMarried := false
	height := 165.1
	nickname := "Fian"
	fmt.Println(name)
	fmt.Println(name, "is", age, "years old")
	fmt.Println(name, "earns", salary)
	fmt.Println(name, "married status is", isMarried)
	fmt.Println(name, "height is", height, "cm")
	fmt.Println(name, "nickname is", nickname)
}

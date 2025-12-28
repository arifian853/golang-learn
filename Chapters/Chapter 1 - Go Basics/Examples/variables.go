package main

import "fmt"

// global variable tanpa identifikasi tipe data (mirip python) dan hanya bisa diluar function
var age = 24 // Integer
var salary = 150000 // Integer
var isMarried = false // Boolean
var height = 165.1 // Float
var nickname = "Fian" // String

// contoh global variable dengan identifikasi tipe data
var hobby string = "Coding, Gaming, Traveling, Sleeping"
var isSleeping bool = true
var clock float64 = 24.0
var phoneNumber int = 123456789

func variables() {
	fmt.Println("\n -------------------")
	fmt.Println("Variables with explicit type")
	var name string = "Arifian Saputra" // local variable, harus identifikasi tipe data dan hanya bisa di dalam function
	fmt.Println(name) 
	fmt.Println(name, "is", age, "years old")
	fmt.Println(name, "earns", salary)
	fmt.Println(name, "married status is", isMarried)
	fmt.Println(name, "height is", height, "cm")
	fmt.Println(name, "nickname is", nickname)
	fmt.Println(name, "hobby is", hobby)
	fmt.Println(name, "is sleeping is", isSleeping)
	fmt.Println(name, "clock is", clock)
	fmt.Println(name, "phone number is", phoneNumber)
}


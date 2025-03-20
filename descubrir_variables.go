package main

import (
	"fmt"
)

func main() {

	c := "cadena de texto"
	i := 8
	f := 4.56789
	fmt.Println("Esto es una línea de texto")
	fmt.Printf("Esto es una %s\n", c)
	fmt.Printf("Hay %d manzanas\n", i)
	fmt.Printf("%f\n", f)

	var myString string = "Mi primer texto"
	fmt.Println(myString)
	var unEntero int = 27
	fmt.Println(unEntero)

	myflotante := 3.5
	fmt.Println(myflotante)

	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age) // Reading input separated by space
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)

}

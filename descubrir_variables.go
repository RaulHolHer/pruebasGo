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

	// Ahora intentaré meter variables con Scan
	var name string
	var age int
	fmt.Print("Por favor dime tu nombre:")
	fmt.Scan(&name) // lee los datos introducidos separados por un espacio
	fmt.Print("Y ahora, por favor, ¿podrías decirme tu edad?: ")
	fmt.Scan(&age)
	fmt.Printf("Hola %s, tienes  %d años.\n", name, age)

}

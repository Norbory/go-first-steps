package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Alex p-....")

	// Variables
	var myString string = "Esto es una cadena"
	fmt.Println(myString)

	myString = "Aqui esta el nuevo valor"
	fmt.Println(myString)

	var myInt int = 77
	fmt.Println(myInt)

	// Printf
	fmt.Printf("%s %d", myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	// Para declarar e inicializar una variable
	myString3 := "Hola"
	fmt.Println(myString3)
}

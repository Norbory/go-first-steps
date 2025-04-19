package main

import (
	"container/list"
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

	var myInt int = 10
	fmt.Println(myInt)

	// Printf
	fmt.Printf("%s %d", myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	// Para declarar e inicializar una variable
	myString3 := "Hola"
	fmt.Println(myString3)

	// Constantes
	// Se pueden declarar y no detener el programa, a pesar de no estarse utilizando
	const myConst string = "Esto es una constante"
	fmt.Println(myConst)

	// Control de flujo
	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

	// Array
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray[2])
	// Si intentas ejecutar uno de estos saldra error
	// myArray[3] = 4
	// fmt.Println(myArray[3])

	// Map
	myMap := make(map[string]int)
	myMap["Brais"] = 36
	myMap["Angelo"] = 35
	myMap["Crais01"] = 24
	fmt.Println(myMap["Brais"])

	myMap2 := map[string]int{"Brais": 25, "Hicomarca": 62, "Hozer&": 64}
	fmt.Println(myMap2["Hicomarca"])

	// List
	// Cada elemento representa un puntero
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Bucles
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	for _, value := range myArray {
		fmt.Println(value)
	}

	// Funciones
	myFunction("Alex")
	fmt.Println(myFunction2("Alex"))
	fmt.Println(myFunction3())

	// Estructuras
	type myStruct struct {
		name string
		age  int
	}

	var myStruct1 = myStruct{name: "Alex", age: 24}
	fmt.Println(myStruct1)

	myStruct2 := myStruct{"Alex", 24}
	fmt.Println(myStruct2)
}

func myFunction(name string) {
	fmt.Println("Hola", name)
}

func myFunction2(name string) string {
	return "Hola " + name
}
func myFunction3() string {
	return "Hola"
}

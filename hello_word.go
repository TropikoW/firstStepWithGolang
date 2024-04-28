package main

import (
	"container/list" // Importa el paquete "list" del paquete "container"
	"fmt"            // Importa el paquete "fmt" para imprimir en la consola
	"reflect"        // Importa el paquete "reflect" para obtener información sobre tipos de datos
)

func main() {
	fmt.Println("Hello, World!") // Imprime "Hello, World!" en la consola

	var myString string = "this is a string text!" // Declara una variable "myString" de tipo string y le asigna un valor
	fmt.Println(myString) // Imprime el valor de "myString" en la consola

	myString = "this is a new string text!" // Reasigna un nuevo valor a la variable "myString"
	fmt.Println(myString) // Imprime el nuevo valor de "myString" en la consola

	var myInt int = 10 // Declara una variable "myInt" de tipo int y le asigna un valor
	fmt.Println(myInt) // Imprime el valor de "myInt" en la consola

	myInt = myInt + 3 // Incrementa el valor de "myInt" en 3
	fmt.Println(myInt) // Imprime el nuevo valor de "myInt" en la consola

	myInt += 3 // Incrementa el valor de "myInt" en 3 usando el operador de asignación abreviada
	fmt.Println(myInt) // Imprime el nuevo valor de "myInt" en la consola

	myInt++ // Incrementa el valor de "myInt" en 1
	fmt.Println(myInt) // Imprime el nuevo valor de "myInt" en la consola

	myInt-- // Decrementa el valor de "myInt" en 1
	fmt.Println(myInt) // Imprime el nuevo valor de "myInt" en la consola

	var stringNumber string = "10" // Declara una variable "stringNumber" de tipo string y le asigna un valor
	var myFloat float64 = 10.5 // Declara una variable "myFloat" de tipo float64 y le asigna un valor

	myString3 := "this is a string text!" // Declara una variable "myString3" y le asigna un valor automáticamente
	fmt.Println(myString3) // Imprime el valor de "myString3" en la consola

	const myConst string = "thi is a constant!" // Declara una constante "myConst" de tipo string y le asigna un valor
	fmt.Println(myConst) // Imprime el valor de "myConst" en la consola
	fmt.Println(stringNumber) // Imprime el valor de "stringNumber" en la consola

	fmt.Println(stringNumber, myInt) // Imprime el valor de "stringNumber" y "myInt" en la consola

	fmt.Println(reflect.TypeOf(stringNumber), reflect.TypeOf(myInt)) // Imprime el tipo de datos de "stringNumber" y "myInt" en la consola

	fmt.Println(reflect.TypeOf(myFloat)) // Imprime el tipo de datos de "myFloat" en la consola

	fmt.Println(float64(myInt) + myFloat) // Imprime la suma de "myInt" convertido a float64 y "myFloat" en la consola

	// Bools

	var myBool bool = true // Declara una variable "myBool" de tipo bool y le asigna un valor
	fmt.Println(myBool) // Imprime el valor de "myBool" en la consola	

	if myBool {
		fmt.Println("The condition is true") // Imprime un mensaje si la condición es verdadera
	}

	if !myBool {
		fmt.Println("The condition is false") // Imprime un mensaje si la condición es falsa
	}

	if myInt > 20 {
		fmt.Println("my int is greater than 20", myInt) // Imprime un mensaje si "myInt" es mayor que 20
	} else {
		fmt.Println("my ins is less than 20") // Imprime un mensaje si "myInt" es menor o igual que 20
	}

	conditional1 := "hola" // Declara una variable "conditional1" y le asigna un valor automáticamente
	contitional2 := "qué tal" // Declara una variable "contitional2" y le asigna un valor automáticamente

	if conditional1 == "hola" && contitional2 == "qué tal" {
		fmt.Println("This conversations is in spanish") // Imprime un mensaje si ambas condiciones son verdaderas
	} else {
		fmt.Println("This conversations is not in spanish") // Imprime un mensaje si alguna de las condiciones es falsa
	}

	// Arrays

	myArray := [3]int{1, 2, 3} // Declara un array "myArray" de tamaño 3 y le asigna valores
	fmt.Println(myArray) // Imprime el contenido de "myArray" en la consola

	var myArray2 [3]int // Declara un array "myArray2" de tamaño 3
	myArray2[0] = 1 // Asigna un valor al primer elemento de "myArray2"
	myArray2[1] = 2 // Asigna un valor al segundo elemento de "myArray2"
	fmt.Println(myArray2) // Imprime el contenido de "myArray2" en la consola

	// Objects

	myMap := make(map[string]int) // Crea un mapa "myMap" con claves de tipo string y valores de tipo int
	myMap["myAge"] = 24 // Asigna un valor al mapa "myMap" con la clave "myAge"

	myMap2 := make(map[string]int) // Crea un mapa "myMap2" con claves de tipo string y valores de tipo int
	myMap2["string"] = 0 // Asigna un valor al mapa "myMap2" con la clave "string"

	fmt.Println(myMap) // Imprime el contenido de "myMap" en la consola
	fmt.Println(myMap2) // Imprime el contenido de "myMap2" en la consola

	myMap3 := map[string]int{"myAge": 24, "oldAge": 23} // Crea un mapa "myMap3" con claves de tipo string y valores de tipo int y les asigna valores
	fmt.Println(myMap3) // Imprime el contenido de "myMap3" en la consola

	// Lists

	myList := list.New() // Crea una lista enlazada vacía "myList"
	fmt.Println("myList:", myList) // Imprime el contenido de "myList" en la consola

	myList.PushBack(1) // Agrega un elemento al final de "myList"
	myList.PushBack(2) // Agrega un elemento al final de "myList"
	myList.PushBack(3) // Agrega un elemento al final de "myList"

	fmt.Println("myList:", myList) // Imprime el contenido de "myList" en la consola
	fmt.Println("myList:", myList.Back().Value) // Imprime el último elemento de "myList" en la consola

	for i := 0; i < 5; i++ {
		fmt.Println(i) // Imprime el valor de "i" en cada iteración del bucle
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i]) // Imprime el valor de cada elemento de "myArray" en cada iteración del bucle
	}

	for key, value := range myMap3 {
		fmt.Println(key, value) // Imprime la clave y el valor de cada elemento de "myMap3" en cada iteración del bucle
	}

	for i, v := range myArray {
		fmt.Println(i, v) // Imprime el índice y el valor de cada elemento de "myArray" en cada iteración del bucle
	}

	myFunctions() // Llama a la función "myFunctions"

	fmt.Println(myFunctionsWithReturn()) // Imprime el valor de retorno de la función "myFunctionsWithReturn"

	// class,in the case of go is a struct

	type User struct { // Define una estructura llamada "User"
		name string // Campo "name" de tipo string en la estructura "User"
		age  int    // Campo "age" de tipo int en la estructura "User"
	}

	var myUser User = User{name: "Jhon", age: 24} // Crea una variable "myUser" de tipo "User" y le asigna valores a los campos
	secondUser := User{name: "Jhon", age: 24} // Crea una variable "secondUser" de tipo "User" y le asigna valores a los campos

	fmt.Println(myUser) // Imprime el contenido de "myUser" en la consola
	fmt.Println(secondUser) // Imprime el contenido de "secondUser" en la consola
}

// functions

func myFunctions() {
	fmt.Println("This is a function") // Imprime un mensaje en la consola
}

func myFunctionsWithReturn() string {
	return "This is a function with return" // Retorna un mensaje
}

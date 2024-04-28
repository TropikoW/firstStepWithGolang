package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// func main() {
// 	print("Hello, World!") <- it's available but's not semantic
// }

func main() {
	fmt.Println("Hello, World!")
	var myString string = "this is a string text!" // var name type = value

	fmt.Println(myString)

	myString = "this is a new string text!" // reassigning a value to a variable

	fmt.Println(myString) // it's semantic, fmt is a package that contains functions for formatting text

	var myInt int = 10

	fmt.Println(myInt)

	myInt = myInt + 3

	fmt.Println(myInt)

	myInt += 3

	fmt.Println(myInt)

	myInt++

	fmt.Println(myInt)

	myInt--

	fmt.Println(myInt)

	var stringNumber string = "10"

	var myFloat float64 = 10.5

	fmt.Println(stringNumber)

	// fmt.Println(stringNumber + string(myInt))

	fmt.Println(stringNumber,myInt)

	fmt.Println(reflect.TypeOf(stringNumber),reflect.TypeOf(myInt))

	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat)

	var myBool bool = true

	fmt.Println(myBool)

	myString3 := "this is a string text!" // var name = value

	fmt.Println(myString3)

	const myConst string = "thi is a constant!"

	fmt.Println(myConst)

	if myBool {
		fmt.Println("The condition is true")
	}

	if !myBool {
		fmt.Println("The condition is false")
	}

	if myInt > 20 {
		fmt.Println("my int is greater than 20",myInt)
	} else {
		fmt.Println("my ins is less than 20")
	}

	conditional1 := "hola"
	contitional2 := "qué tal"

	if conditional1 == "hola" && contitional2 == "qué tal" {
		fmt.Println("This conversations is in spanish")	
	} else {
		fmt.Println("This conversations is not in spanish")
	}

	myArray := [3]int{1,2,3}

	fmt.Println(myArray)


	var myArray2 [3]int

	myArray2[0] = 1

	myArray2[1] = 2

	fmt.Println(myArray2)

	myMap := make(map[string]int)

	myMap["myAge"] = 24

	myMap2 := make(map[string]int) // map[keyType]valueType

	myMap2["string"] = 0

	fmt.Println(myMap)

	fmt.Println(myMap2)

	myMap3 := map[string]int{"myAge": 24,"oldAge":23} // map[keyType]valueType

	fmt.Println(myMap3)

	myList := list.New()

	fmt.Println("myList:",myList)

	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	fmt.Println("myList:",myList)

	fmt.Println("myList:",myList.Back().Value)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for key /*it's represent the index*/, value := range myMap3 {
		fmt.Println(key,value)
	}

	for i,v := range myArray {
		fmt.Println(i,v)	
	}

	myFunctions()

	fmt.Println(myFunctionsWithReturn())

	// structures

	type User struct {
		name string
		age int
	}

	var myUser User = User{name: "Jhon", age: 24}

	secondUser := User{name: "Jhon", age: 24}

	fmt.Println(myUser)

	fmt.Println(secondUser)
}

func myFunctions(){
	fmt.Println("This is a function")
}

func myFunctionsWithReturn() string{
	return "This is a function with return"	
}
// go run hello_word.go  // or go run . (to run actual page) // or go run * or go run main.go

// go build hello_word.go
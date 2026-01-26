package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello World")

	// integer
	var intNum uint = 32767 // uint is unsigned int | int is signed int
	// uint8, uint16, uint32, uint64
	// int8, int16, int32, int64
	intNum = intNum + 10
	fmt.Println(intNum)

	// float
	var floatNum float64 = 3.14 // float32, float64  | float32 --> 6 digits | float64 --> 15 digits
	fmt.Println(floatNum)

	// string
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("Y"))                       // len() is used to get the length of a string | this prints no of bytes
	fmt.Println(utf8.RuneCountInString("Yash")) // utf8.RuneCountInString() is used to get the length of a string | this prints no of characters

	//rune (char)
	var myRune rune = 'a'
	fmt.Println(myRune)

	//boolean
	var myBool bool = true
	fmt.Println(myBool)

	// default value check | default values for int,float,rune = 0 | default value for bool = false | string = ""
	var intNum3 int
	fmt.Println(intNum3)

	var myVar = "text"
	fmt.Println(myVar)

	myVar2 := "text"
	fmt.Println(myVar2)

	var1, var2 := "text", "text2"
	fmt.Println(var1, var2)

	var3, var4 := 1, 2
	fmt.Println(var3, var4)

	// const values
	const myConst string = "my const" // const values cannot be changed
	fmt.Println(myConst)

}

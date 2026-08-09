package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {

	/*
		SIGNED INTEGERS
		int     - -2147483648 to 2147483647 (32 bit system, 2^32 total values) / -9223372036854775808 to 9223372036854775807 (64 bit system, 2^64 total values)
		int8    - -128 to 127 (256 total values)
		int16   - -32768 to 32767 (2^16 total values)
		int32   - -2147483648 to 2147483647 (2^32 total values)
		int64   - -9223372036854775808 to 9223372036854775807 (2^64 total values)

		UNSIGNED INTEGERS
		uint	- 0 to 4294967295 (32 bit system) / 0 to 18446744073709551615 (64 bit system)
		uint8	- 0 to 255
		uint16	- 0 to 65535
		uint32	- 0 to 4294967295
		uint64	- 0 to 18446744073709551615

		default value is 0 for all
	*/
	var intNum int = 100
	fmt.Println(intNum)

	/*
		float32				- (32 bits) - -3.4e+38 to 3.4e+38.
		float64 (default)	- (64 bits) - -1.7e+308 to +1.7e+308.

		default value is 0 for both
	*/
	var floatNum float64 = 100.56
	fmt.Println(floatNum)

	/*
		STRING
		default value for string is ''
	*/
	fmt.Println("Hello World!")

	// String length
	fmt.Println(len("A@")) // this returns the number of bytes. not the number of characters.

	fmt.Println(utf8.RuneCountInString("A@")) // rune: data type for characters

	/*
		BOOLEAN

		default value is false
	*/
	var myBoolean bool = false
	fmt.Println(myBoolean)

	const myConst string = "const value" // cannot be changed later and should always have a value assigned

	// =================================================

	// FUNCTIONS & CONSTROL STRUCTURES
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 1
	var denominator int = 1
	var result, remainder, err = intDivision(numerator, denominator)

	// handle using if else statements
	fmt.Println("Handle with if else")
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v \n", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	// handle using switch
	fmt.Println("Handle with switch")
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v \n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	// handle using switch with a variable
	fmt.Println("Handle with switch with a variable")
	switch remainder {
	case 0:
		fmt.Printf("The result of the integer division is %v \n and the remainder is 0 \n", result)
	case 1, 2:
		fmt.Printf("The result of the integer division is %v \n and the remainder is 1 or 2 \n", result)
	default:
		fmt.Printf("The result of the integer division is %v \n and the remainder is %v \n", result, remainder)
	}

	fmt.Printf("The result of the integer divisin is %v with remainder %v \n", result, remainder)

	// ==============================================================
	// ARRAYS, SLICES, MAPS AND LOOPING

	// ARRAYS
	var intArr [3]int32      // 4 bytes x 3 = 12 bytes
	fmt.Println(intArr)      // prints [0 0 0]
	fmt.Println(intArr[0])   // prints 0
	fmt.Println(intArr[1:3]) // prints [0 0]

	intArr[0] = 123
	fmt.Println(intArr) // prints [123 0 0]

	// print memory address
	fmt.Println(&intArr[0]) // prints memory address
	fmt.Println(&intArr[1]) // prints memory address
	fmt.Println(&intArr[2]) // prints memory address
	// memory location of array is contiguous

	// initialize an array
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// or
	intArr3 := [3]int32{1, 2, 3} // or [...]int32{1, 2, 3}
	fmt.Println(intArr3)

	// SLICES
	var intSlice []int32 = []int32{1, 2, 3, 5}
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// note: slices look for the available space, if not it allocate the same number of batch positions again
	// ex: intSlice has only 4 indexes at the first place but it looks for an empty space to add 7
	// then it allocate another 4 spaces. finally the capacity is 8. ----> [1 2 3 5 7 * * *]. but cant access additional indexes

	// append another slice
	var exSlice []int32 = []int32{6, 7, 8}
	intSlice = append(intSlice, exSlice...)
	fmt.Println(intSlice)

	// initialize a slice
	var intSlice3 []int32 = make([]int32, 3, 8) // length 3 , capacity 8. if capacity isnt provided, capacity would be the length of the slice
	fmt.Println(intSlice3)
	fmt.Printf("Length %v and Capacity %v \n", len(intSlice3), cap(intSlice3))

	// MAP
	var myMap1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap1)

	var myMap2 = map[string]uint8{"Adam": 48, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])  // prints 48
	fmt.Println(myMap2["Jason"]) // prints 0

	// how to omit default value if the key doesnt exists
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v \n", age)
	} else {
		fmt.Printf("Invalid name \n")
	}

	delete(myMap2, "Adam") // delete a key with value from a map
	fmt.Println(myMap2)

	// ITERATE over MAP , ARRAY or SLICE

	// over map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	// over array/slices
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// GO doesnt have a while loop. but we can create using for
	var i int = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	// or
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// understand the loop over slices
	var n int = 1000000
	var testSlice = []int{}            // without pre allocation
	var testSlice2 = make([]int, 0, n) // with preallocation

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))

	// this prints (these values can be changed) --> this says preallocation increases the performance
	//Total time without preallocation: 22.837445ms
	//Total time with preallocation: 5.187835ms

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Devide by Zero")
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

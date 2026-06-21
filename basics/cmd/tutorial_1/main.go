package main

import "fmt"

func main() {

	// print hello word
	fmt.Println("Hello World!")

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
	*/
	var intNum int
	fmt.Println(intNum)
}

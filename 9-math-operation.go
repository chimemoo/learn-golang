/**
 *
 * Math Operation in Go
 *
 * In Go, we can only use operators on the same
 * data types. Example we can'1 add an int and a float64
 * you must convert one of that data type
 *
 */
package main

import "fmt"

func main() {

	var result = 17 + 13
	fmt.Println(result)

	var a = 1
	var b float32 = 9.5
	fmt.Println(float32(a) + b)

}

/**
 *
 * Data type convertion in Go
 *
 * Be aware to using data type convertion, example :
 * if use int32 and convert to smaller data type like int16
 * thats will change the value and make overflow integer because
 * the new data type cannot
 * receive the int32 value
 *
 */

package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var firstName = "Eko"
	var e = firstName[0]    // Get byte of E
	var eString = string(e) // Convert byte to string

	fmt.Println(firstName)
	fmt.Println(e)
	fmt.Println(eString)
}

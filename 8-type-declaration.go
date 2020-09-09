/**
 *
 * Type Declarations in Go
 *
 * Type declaration is ability to make new data type from
 * available data type
 *
 *
 */

package main

import "fmt"

func main() {
	type IdNumber string // Create new data type/alias

	var userIdNumber IdNumber = "19929919"
	fmt.Println(userIdNumber)
}

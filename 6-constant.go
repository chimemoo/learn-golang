/**
 *
 * Constant in Go
 *
 * Constant is variable that value never can change
 * Key to declare constant is const
 * You must pass value if declare the constant variable
 * Not like var, you can not use constant variable if you declare it.
 *
 */

package main

import "fmt"

func main() {
	const firstName string = "Asep"
	const lastName string = "Saepudin"

	fmt.Println(firstName + " " + lastName)
}

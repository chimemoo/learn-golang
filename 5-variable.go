/**
 *
 * Variable in Go
 * Key for define variable is using var following by variable name
 * and data type
 *
 * If you declare variable, you must using it!, because go compiler
 * will showing error. Even you pass value to variable and
 * not using that variable go compiler will showing error
 *
 * You can declare variable without define the data type because
 * go compiler automatically detect the data type
 *
 * You can define variable without "var" but you must declare value of variable
 * using ":=" on first declaration
 *
 * To declare multiple variable using var ()
 *
 */

package main

import "fmt"

func main() {
	var name string
	var id int
	var companyId = 1
	companyName := "Github"

	var (
		country = "Indonesia"
	)

	name = "Chimemoo"
	id = 1
	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(companyId)
	fmt.Println("Company Name : " + companyName)
	fmt.Println("Country " + country)

	name = "Github"
	id = 2
	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(companyId)
	fmt.Println("Company Name : " + companyName)
	fmt.Println("Country " + country)

}

/**
 *
 * Hello World Program in GO
 *
 * What i learn :
 * - fmt is package for showing text
 * - to run go you must compile go file using "go build namefile.go",
 *   after compile "go" will make new file without extension
 * 	 that can run in system with the same os
 * - To run the compiled file, in linux only run with "./namefile"
 * - Go file can run without compile with "go run namefile.go" (Only in development)
 * - In production only upload the binary file
 *
 */

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

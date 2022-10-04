// Modify the echo program to print the index and value of each of its arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	i := 1
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

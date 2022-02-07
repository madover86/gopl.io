package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Argument #%v is %s\n", i, arg)
	}
}
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main ()  {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args {
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
	diff1 := time.Now().Nanosecond() - start.Nanosecond()
	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	diff2 := time.Now().Nanosecond() - start.Nanosecond()
	fmt.Printf("Result 1 = %v\nResult 2 = %v", diff1, diff2)
}
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep, in := "", "", ""
	for i, arg := range os.Args[0:] {
		in = strconv.Itoa(i)
		sep = "\n"
		s += in + arg + sep
	}
	fmt.Println(s)
}

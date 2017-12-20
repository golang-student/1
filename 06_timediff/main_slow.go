package main

import (
	"fmt"
	"os"
	"time"
	//"strings"
)

func main() {
	var s, sep string
	t := time.Now()	
	for _, arg := range os.Args[1:] {
		s = s + sep + arg
		sep = " "
	}
	fmt.Println(s)
	diff := time.Now().Sub(t)
	fmt.Println("elapsed time = " + diff.String())
}

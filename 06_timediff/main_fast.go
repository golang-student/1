package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	t := time.Now()	
	fmt.Println(strings.Join(os.Args[1:]," "))
	diff := time.Now().Sub(t)
	fmt.Println("elapsed time = " + diff.String())
}

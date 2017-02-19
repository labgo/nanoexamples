package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
}

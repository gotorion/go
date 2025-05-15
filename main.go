package main

import(
	"fmt"
	"runtime"
	"github.com/gotorion/go/algorithm"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(runtime.Version())
	fmt.Println(algorithm.Add(1, 2))
}
package main

import (
    "fmt"
    "runtime"
	"err"
)

func main() {
	foo()
}

func foo() {
    err.Err("some error")
}



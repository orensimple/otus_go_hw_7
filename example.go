package main

import (
	"fmt"
	"os"
)

func main() {
	var1, _ := os.LookupEnv("A_ENV")
	fmt.Printf(var1) // val
	var2, _ := os.LookupEnv("B_VAR")
	fmt.Printf("   %s", var2) // val
}

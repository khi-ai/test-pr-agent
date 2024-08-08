package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	fmt.Println(" 1 + 1 = ", 1+1)

	fmt.Printf("Return sample: %s\n", sample())
}

func sample() string {
	return "sample"
}

package main

import "fmt"

func main() {
	text := "This is a header <-<Header-"

	a, err := parseArrow(text)

	fmt.Println(a, err)
}

package main

import "fmt"

func main() {
	text := "This is a header <-<Header- shit"

	nodes, _ := parseArrow(text)

	fmt.Println(nodes)
}

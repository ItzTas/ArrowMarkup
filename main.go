package main

import "fmt"

func main() {
	text := `This is a header <-Header-`
	parser := NewAmParser()

	fmt.Println(parser.parseAM(text))

}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := ".class.< testclass anotherclass > .href.< testLink >"
	regex := regexp.MustCompile(`\.(\w+)\.<([^>]+)>`)

	matches := regex.FindAllString(text, -1)
	fmt.Println(parseAttributes(matches))
}

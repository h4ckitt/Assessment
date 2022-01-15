package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	numReg  = regexp.MustCompile("[0-9-]+")
	textReg = regexp.MustCompile("[a-zA-Z-]+")
)

func testValidity(text string) bool {
	split := strings.Split(text, "-")

	if len(split) < 2 {
		return false
	}

	allTextReg := regexp.MustCompile("[a-zA-Z]+")
	allNumReg := regexp.MustCompile("[0-9]+")

	for _, elem := range split {
		if !allNumReg.MatchString(elem) && !allTextReg.MatchString(elem) {
			return false
		}
	}
	return true
}

func main() {
	text := "23-ab-48-caba-56-haha-2"
	fmt.Println(testValidity(text))
}

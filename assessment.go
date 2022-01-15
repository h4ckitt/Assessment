package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	allNumReg  = regexp.MustCompile("[0-9]+")
	allTextReg = regexp.MustCompile("[a-zA-Z+")
)

func testValidity(text string) bool {
	split := strings.Split(text, "-")

	if len(split) < 2 {
		return false
	}

	//	allTextReg := regexp.MustCompile("[a-zA-Z]+")
	//	allNumReg := regexp.MustCompile("[0-9]+")

	for _, elem := range split {
		if !allNumReg.MatchString(elem) && !allTextReg.MatchString(elem) {
			return false
		}
	}
	return true
}

func averageNumber(text string) float64 {
	split := strings.Split(text, "-")

	var (
		total    int
		numCount int
	)

	for _, elem := range split {
		if allNumReg.MatchString(elem) {
			num, _ := strconv.Atoi(elem)
			total += num
			numCount++
		}
	}

	average := float64(total) / float64(numCount)

	return float64(int(average*100)) / 100
}

func main() {
	text := "23-ab-48-caba-56-haha-2"
	fmt.Println(testValidity(text))
}

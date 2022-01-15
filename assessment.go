package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	allNumReg  = regexp.MustCompile("[0-9]+")
	allTextReg = regexp.MustCompile("[a-zA-Z+]")
)

/* testValidity:
   Difficulty: Easy
   Used Time Of Completion: 10 Minutes
*/
func testValidity(text string) bool {
	split := strings.Split(text, "-")

	if len(split) < 2 {
		return false
	}

	for _, elem := range split {
		if !allNumReg.MatchString(elem) && !allTextReg.MatchString(elem) {
			return false
		}
	}
	return true
}

/* averageumber:
Difficulty: Easy
Used Time Of Completion: 10 Minutes
*/

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

/* wholeStory:
Difficulty: Easy
Used Time Of Completion: 8 Minutes
*/

func wholeStory(text string) string {
	var result []string

	split := strings.Split(text, "-")

	for _, elem := range split {
		if allTextReg.MatchString(elem) {
			result = append(result, elem)
		}
	}

	return strings.Join(result, " ")
}

/* storyStats:
Difficulty: Easy
Used Time Of Completion: 13 Minutes
*/

func storyStats(text string) (string, string, float64, []string) {
	split := strings.Split(text, "-")
	var (
		words           []string
		longestWord     string
		shortestWord    string
		resultList      []string
		totalWordLength int
	)

	for _, elem := range split {
		if allTextReg.MatchString(elem) {
			words = append(words, elem)
		}
	}

	shortestWord = words[0]

	for _, elem := range words {
		length := len(elem)

		if length < len(shortestWord) {
			shortestWord = elem
		} else if length > len(longestWord) {
			longestWord = elem
		}

		totalWordLength += length
	}

	average := float64(totalWordLength) / float64(len(words))
	roundedAverage := int(math.Round(average))

	for _, elem := range words {
		if len(elem) == roundedAverage {
			resultList = append(resultList, elem)
		}
	}

	return shortestWord, longestWord, float64(int(average*100)) / 100, resultList

}

func main() {
	text := "23-ab-48-caba-56-haha-2"
	fmt.Println(testValidity(text))
	fmt.Println(averageNumber(text))
	fmt.Println(wholeStory(text))
	fmt.Println(storyStats(text))
}

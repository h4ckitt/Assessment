package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
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

/* generate:
Difficulty: Medium
Estimated Time Of Completion: 20 Minutes
Used Time Of Completion: 50 Minutes
*/

func generate(valid bool) string {
	var result strings.Builder
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(12-2) + 2
	if valid {
		for i := 0; i <= length; i++ {
			if i%2 == 0 {
				num := rand.Intn(10000-1) + 1
				result.WriteString(strconv.Itoa(num))
				continue
			}
			result.WriteString("-" + randString(rand.Intn(5-1)+1, true) + "-")
		}
		return strings.TrimLeft(strings.TrimRight(result.String(), "-"), "-")
	}

	for i := 0; i <= length; i++ {
		result.WriteString(randString(rand.Intn(5-1)+1, false))
	}

	return result.String()
}

func randString(length int, valid bool) string {
	var letterBytes string
	if valid {
		letterBytes = "abcdefghijklmnopqrstuvwxyz" //ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890-"
	}

	b := make([]byte, length)

	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}

func main() {
	text := "1466-boi-3784-guw-7113-rqgw-926-bht-8681-e-7143"
	fmt.Println(testValidity(text))
	fmt.Println(averageNumber(text))
	fmt.Println(wholeStory(text))
	fmt.Println(storyStats(text))
	fmt.Println(generate(true))
	fmt.Println(generate(false))
}

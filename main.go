package main

import (
	"assessment/assessment"
	"fmt"
)

func main() {
	//	text := "1466-boi-3784-guw-7113-rqgw-926-bht-8681-e-7143"
	text := "20-ab-50-cd-99-df"
	fmt.Println(assessment.TestValidity(text))
	fmt.Println(assessment.StoryStats(text))
	/*        fmt.Println(averageNumber(text))
	          fmt.Println(wholeStory(text))
	          fmt.Println(storyStats(text))
	          fmt.Println(generate(true))
	          fmt.Println(generate(false))
	*/
}

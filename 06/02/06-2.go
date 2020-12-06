package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func toBinary(answer string) int64 {
	// for each set of answers, create a binary representation of the
	// answers. (i.e. for if you 5 questions: "a", "b", "c", "d", and "e"
	// and you answer "yes" to "a" and "d", your binary number would
	// be '10010')
	bin := []rune(strings.Repeat("0", 26))
	for _, char := range strings.ToLower(answer) {
		// the lower-case ASCII letters begin at 97, so we
		// need to shift our characters down by 97 to get the
		// correct index
		//
		// http://www.asciitable.com/
		index := char - 97
		bin[index] = '1'
	}

	ans, _ := strconv.ParseInt(string(bin), 2, 64)
	return ans
}

func main() {
	// get the "answer" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	answerText := *input

	// split the answer text into groups and convert them
	// all to binary
	answerGroups := strings.Split(answerText, "\n\n")
	allAnswersBin := [][]int64{}
	for _, groupAnswers := range answerGroups {
		individualAnswers := strings.Split(groupAnswers, "\n")
		groupAnswersBin := []int64{}
		for _, answers := range individualAnswers {
			bin := toBinary(answers)
			groupAnswersBin = append(groupAnswersBin, bin)
		}
		allAnswersBin = append(allAnswersBin, groupAnswersBin)
	}

	// for each group of answers, find all commonalities between
	// them using the 'bitwise AND' operation
	commonAnswerCount := 0
	for _, groupAnswers := range allAnswersBin {
		// this creates a series of 26 "1"s in binary
		commonAnswers, _ := strconv.ParseInt(string([]rune(strings.Repeat("1", 26))), 2, 64)
		// iterate through the list, applying a series of
		// "bitwise AND"s
		for _, answers := range groupAnswers {
			commonAnswers = commonAnswers & answers
		}
		// count the number of "1"s in the list
		commonAnswerCount += strings.Count(strconv.FormatInt(commonAnswers, 2), "1")
	}
	fmt.Println(commonAnswerCount)
}

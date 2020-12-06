package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

/**
 * Given a string of answers, create a lookup map, where each answer
 * corresponds to a key.
 */
func makeAnswerLog(answers string) map[string]int {
	answerMap := make(map[string]int)
	for _, char := range answers {
		answerMap[string(char)] = 1
	}
	return answerMap
}

func main() {
	// get the "answer" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}
	answerText := *input

	// split the answer text into groups and flatten
	// the answers
	answerGroups := strings.Split(answerText, "\n\n")
	flattenedAnswerGroups := []string{}
	for _, answers := range answerGroups {
		flattendAnswers := strings.ReplaceAll(answers, "\n", "")
		flattenedAnswerGroups = append(flattenedAnswerGroups, flattendAnswers)
	}

	/**
	 * create an answer map for each of our characters
	 *
	 * @NOTE - instead of creating a map to store this
	 * data, we could instead just flatten our string
	 * (i.e. remove duplicate) entries, then take the
	 * length. It would probably be faster, but I prefer
	 * this data structure, since it's more flexible
	 *
	 * ~reccanti 12/6/2020
	 */
	logs := []map[string]int{}
	for _, answers := range flattenedAnswerGroups {
		log := makeAnswerLog(answers)
		logs = append(logs, log)
	}

	// get the sum of all of our answers
	sum := 0
	for _, log := range logs {
		// get the sum of the answer in each group
		groupSum := 0
		for _, val := range log {
			groupSum += val
		}
		sum += groupSum
	}

	// print the result
	fmt.Println(sum)

}

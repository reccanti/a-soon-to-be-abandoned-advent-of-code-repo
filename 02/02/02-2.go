package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	char string
	low  int
	high int
}

type Password struct {
	value  string
	policy Policy
}

// A utility function for parsing an input file into an
// array of values.
//
// @TODO If this is a common problem later, it might make
// sense to break this down into some utility functions
// that can be reused in future problems
//
// ~reccanti 12/2/2020
func parseFile(filename string) ([]Password, error) {
	// 1. get the file name
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filepath := path.Join(wd, filename)

	// 2. get the string contents of the file
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// 3. Parse that string into the Password struct
	exp := regexp.MustCompile(`(\d*)\-(\d*) (.): (.*)`)
	strs := strings.Split(string(dat), "\n")
	passwords := []Password{}
	for _, str := range strs {
		// get values from string
		res := exp.FindStringSubmatch(str)[1:]
		lowStr := res[0]
		highStr := res[1]
		char := res[2]
		value := res[3]

		low, err := strconv.Atoi(lowStr)
		if err != nil {
			return nil, err
		}
		high, err := strconv.Atoi(highStr)
		if err != nil {
			return nil, err
		}

		// construct password
		pass := Password{
			value: value,
			policy: Policy{
				low:  low,
				high: high,
				char: char,
			},
		}

		// add that to password array
		passwords = append(passwords, pass)
	}

	return passwords, nil
}

func checkPasswordPolicy(password Password) bool {
	// Check to make sure that the password has at
	// least as many characters as the highest character
	// we want to check against
	if len(password.value) < password.policy.high {
		return false
	}

	// check for character matches at the given position in the password string
	lowMatch := string(password.value[password.policy.low-1]) == password.policy.char
	highMatch := string(password.value[password.policy.high-1]) == password.policy.char

	// this is XOR so return false if BOTH the low and high
	// values match
	if lowMatch && highMatch {
		return false
	} else if lowMatch || highMatch {
		return true
	} else {
		return false
	}
}

func main() {
	filename := os.Args[1]

	passwords, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	matches := 0
	for _, pass := range passwords {
		doesPass := checkPasswordPolicy(pass)
		if doesPass {
			matches += 1
		}
	}
	fmt.Println(matches)
}

package util

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func ParseRelativeFile(filename string) (*string, error) {
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
	str := string(dat)

	// 3. Return that string
	return &str, nil
}

func ParseRelativeFileSplit(filename string, separator string) ([]string, error) {
	input, err := ParseRelativeFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(*input, separator), nil
}

func ParseRelativeFileInts(filename string, separator string) ([]int, error) {
	inputs, err := ParseRelativeFileSplit(filename, separator)
	if err != nil {
		return nil, err
	}
	nums := []int{}
	for _, val := range inputs {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

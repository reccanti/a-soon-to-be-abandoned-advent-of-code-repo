package util

import (
	"io/ioutil"
	"os"
	"path"
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

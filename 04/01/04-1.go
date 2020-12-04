package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func main() {
	// get the "hill" data
	filename := os.Args[1]
	input, err := util.ParseRelativeFile(filename)
	if err != nil {
		return
	}

	// passports are separated by a blank line, so split the string by
	// two consecutive newline characters
	passportBlocks := strings.Split(*input, "\n\n")

	// for each passport block, parse its values into a map of
	// passport fields
	exp := regexp.MustCompile(`\s`)
	passports := []map[string]string{}
	for _, block := range passportBlocks {
		fields := exp.Split(block, -1)
		passport := make(map[string]string)
		for _, field := range fields {
			fieldVals := strings.Split(field, ":")
			passport[fieldVals[0]] = fieldVals[1]
		}
		passports = append(passports, passport)
	}

	// validate the passport
	validPassports := []map[string]string{}
	for _, passport := range passports {

		// Potential Fields:
		//
		// byr (Birth Year)
		// iyr (Issue Year)
		// eyr (Expiration Year)
		// hgt (Height)
		// hcl (Hair Color)
		// ecl (Eye Color)
		// pid (Passport ID)
		// cid (Country ID)

		_, hasByr := passport["byr"]
		_, hasIyr := passport["iyr"]
		_, hasEyr := passport["eyr"]
		_, hasHgt := passport["hgt"]
		_, hasHcl := passport["hcl"]
		_, hasEcl := passport["ecl"]
		_, hasPid := passport["pid"]

		if hasByr && hasIyr && hasEyr && hasHgt && hasHcl && hasEcl && hasPid {
			validPassports = append(validPassports, passport)
		}
	}

	fmt.Println(len(validPassports))
}

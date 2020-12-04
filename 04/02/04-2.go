package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reccanti/a-soon-to-be-abandoned-advent-of-code-repo/util"
)

func isValidByr(byr string) bool {
	// make sure the string is exactly 4 characters
	if len(byr) != 4 {
		return false
	}
	// make sure the string is an integer
	i, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	// make sure the date is between 1920 and 2002
	if i >= 1920 && i <= 2002 {
		return true
	} else {
		return false
	}
}

func isValidIyr(iyr string) bool {
	// make sure the string is exactly 4 characters
	if len(iyr) != 4 {
		return false
	}
	// make sure the string is an integer
	i, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	// make sure the date is between 2010 and 2020
	if i >= 2010 && i <= 2020 {
		return true
	} else {
		return false
	}
}

func isValidEyr(eyr string) bool {
	// make sure the string is exactly 4 characters
	if len(eyr) != 4 {
		return false
	}
	// make sure the string is an integer
	i, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	// make sure the date is between 2020 and 2030
	if i >= 2020 && i <= 2030 {
		return true
	} else {
		return false
	}
}

func isValidHgt(hgt string) bool {
	// attempt to parse the string
	exp := regexp.MustCompile(`(\d*)(in|cm)`)
	res := exp.FindStringSubmatch(hgt)
	// make sure the string matches the structure we're looking for
	if len(res) < 3 {
		return false
	}
	// validate the height ranges for 'in' and 'cm' height units
	if res[2] == "cm" {
		// based on our regex, we know this can be converted to an int
		cms, _ := strconv.Atoi(res[1])
		if cms >= 150 && cms <= 193 {
			return true
		}
	} else if res[2] == "in" {
		// based on our regex, we know this can be converted to an int
		ins, _ := strconv.Atoi(res[1])
		if ins >= 59 && ins <= 76 {
			return true
		}
	}
	return false
}

func isValidHcl(hcl string) bool {
	exp := regexp.MustCompile(`#[a-f0-9]{6}$`)
	res := exp.FindStringSubmatch(hcl)
	if len(res) < 1 {
		return false
	}
	return true
}

func isValidEcl(ecl string) bool {
	if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
		return true
	}
	return false
}

func isValidPid(pid string) bool {
	exp := regexp.MustCompile(`^\d{9}$`)
	res := exp.FindStringSubmatch(pid)
	if len(res) < 1 {
		return false
	}
	return true
}

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

		isValid := true

		byr, hasByr := passport["byr"]
		if !(hasByr && isValidByr(byr)) {
			isValid = isValid && false
		}
		iyr, hasIyr := passport["iyr"]
		if !(hasIyr && isValidIyr(iyr)) {
			isValid = isValid && false
		}
		eyr, hasEyr := passport["eyr"]
		if !(hasEyr && isValidEyr(eyr)) {
			isValid = isValid && false
		}
		hgt, hasHgt := passport["hgt"]
		if !(hasHgt && isValidHgt(hgt)) {
			isValid = isValid && false
		}
		hcl, hasHcl := passport["hcl"]
		if !(hasHcl && isValidHcl(hcl)) {
			isValid = isValid && false
		}
		ecl, hasEcl := passport["ecl"]
		if !(hasEcl && isValidEcl(ecl)) {
			isValid = isValid && false
		}
		pid, hasPid := passport["pid"]
		if !(hasPid && isValidPid(pid)) {
			isValid = isValid && false
		}

		if isValid {
			validPassports = append(validPassports, passport)
		}
	}

	fmt.Println(len(validPassports))
}

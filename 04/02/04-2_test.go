package main

import (
	"testing"
)

func TestByrValidation(t *testing.T) {
	longStringValid := isValidByr("02000")
	if longStringValid != false {
		t.Errorf("byr should error out if the string is too long")
	}

	nonIntStringValid := isValidByr("test")
	if nonIntStringValid != false {
		t.Errorf("byr should error out if it can't be converted to an int")
	}

	oldDateValid := isValidByr("1900")
	if oldDateValid != false {
		t.Errorf("byr should error out if the date is lower than 1920")
	}

	recentDateValid := isValidByr("2010")
	if recentDateValid != false {
		t.Errorf("byr should error out if the date is greater than 2002")
	}

	validDateValid := isValidByr("2000")
	if validDateValid != true {
		t.Errorf("byr should validate a date between 1920 and 2002")
	}
}

func TestIyrValidation(t *testing.T) {
	longStringValid := isValidIyr("02015")
	if longStringValid != false {
		t.Errorf("iyr should error out if the string is too long")
	}

	nonIntStringValid := isValidIyr("test")
	if nonIntStringValid != false {
		t.Errorf("iyr should error out if it can't be converted to an int")
	}

	oldDateValid := isValidIyr("2000")
	if oldDateValid != false {
		t.Errorf("iyr should error out if the date is lower than 2010")
	}

	recentDateValid := isValidIyr("2030")
	if recentDateValid != false {
		t.Errorf("iyr should error out if the date is greater than 2020")
	}

	validDateValid := isValidIyr("2015")
	if validDateValid != true {
		t.Errorf("iyr should validate a date between 2010 and 2020")
	}
}

func TestEyrValidation(t *testing.T) {
	longStringValid := isValidEyr("02015")
	if longStringValid != false {
		t.Errorf("eyr should error out if the string is too long")
	}

	nonIntStringValid := isValidEyr("test")
	if nonIntStringValid != false {
		t.Errorf("eyr should error out if it can't be converted to an int")
	}

	oldDateValid := isValidEyr("2015")
	if oldDateValid != false {
		t.Errorf("eyr should error out if the date is lower than 2020")
	}

	recentDateValid := isValidEyr("2050")
	if recentDateValid != false {
		t.Errorf("eyr should error out if the date is greater than 2030")
	}

	validDateValid := isValidEyr("2025")
	if validDateValid != true {
		t.Errorf("eyr should validate a date between 2020 and 2030")
	}
}

func TestHgtValidation(t *testing.T) {
	measurementValid := isValidHgt("test")
	if measurementValid != false {
		t.Errorf("hgt should error out if it is not a number followed by a measurement string")
	}

	unitValid := isValidHgt("20tt")
	if unitValid != false {
		t.Errorf("hgt should error out if it's unit is not 'in' or 'cm'")
	}

	lowCmValid := isValidHgt("100cm")
	if lowCmValid != false {
		t.Errorf("hgt should error out if the 'cm' height is less than 150")
	}

	highCmValid := isValidHgt("200cm")
	if highCmValid != false {
		t.Errorf("hgt should error out if the 'cm' height is greater than 193")
	}

	lowInValid := isValidHgt("30in")
	if lowInValid != false {
		t.Errorf("hgt should error out if the 'in' height is less than 59")
	}

	highInValid := isValidHgt("100in")
	if highInValid != false {
		t.Errorf("hgt should error out if the 'in' height is greater than 76")
	}

	validCmValid := isValidHgt("175cm")
	if validCmValid != true {
		t.Errorf("hgt should validate out if the 'cm' height is between 150 and 193")
	}

	validInValid := isValidHgt("65in")
	if validInValid != true {
		t.Errorf("hgt should validate out if the 'in' height is between 59 and 76")
	}
}

func TestHclValidation(t *testing.T) {
	invalidHexValid := isValidHcl("#123456ab")
	if invalidHexValid != false {
		t.Errorf("hcl should error out if the hex code is invalid")
	}

	validHexValid := isValidHcl("#01ab03")
	if validHexValid != true {
		t.Errorf("hcl should validate if the hex code is valid")
	}
}

func TestEclValidation(t *testing.T) {
	ambValid := isValidEcl("amb")
	if ambValid != true {
		t.Errorf("ecl should validate 'amb'")
	}
	bluValid := isValidEcl("blu")
	if bluValid != true {
		t.Errorf("ecl should validate 'blu'")
	}
	brnValid := isValidEcl("brn")
	if brnValid != true {
		t.Errorf("ecl should validate 'brn'")
	}
	gryValid := isValidEcl("gry")
	if gryValid != true {
		t.Errorf("ecl should validate 'gry'")
	}
	grnValid := isValidEcl("grn")
	if grnValid != true {
		t.Errorf("ecl should validate 'grn'")
	}
	hzlValid := isValidEcl("hzl")
	if hzlValid != true {
		t.Errorf("ecl should validate 'hzl'")
	}
	othValid := isValidEcl("oth")
	if othValid != true {
		t.Errorf("ecl should validate 'oth'")
	}
	testValid := isValidEcl("test")
	if testValid != false {
		t.Errorf("ecl should only pass one of these values: amb blu brn gry grn hzl oth")
	}
}

func TestPidValidation(t *testing.T) {
	shortPidValid := isValidPid("12345678")
	if shortPidValid != false {
		t.Errorf("pids less than nine characters should be marked as invalid")
	}

	longPidValid := isValidPid("0123456789")
	if longPidValid != false {
		t.Errorf("pids more than nine characters should be marked as invalid")
	}

	pidValid := isValidPid("123456789")
	if pidValid != true {
		t.Errorf("pid should pass if it's exactly nine characters")
	}
}

package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"regexp"
)
type Passport struct{
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
	length_of_fields int
}
func parseFields(pport string) Passport {
	fields := strings.Fields(pport)
	pass := Passport{}
	pass.length_of_fields = len(fields)
	for _, field := range fields {
		info := strings.Split(field, ":")
		if info[0] == "byr" {
			pass.byr, _ = strconv.Atoi(info[1])
		} else if info[0] == "iyr" {
			pass.iyr, _ = strconv.Atoi(info[1])
		} else if info[0] == "eyr" {
			pass.eyr, _ = strconv.Atoi(info[1])
		} else if info[0] == "hgt" {
			pass.hgt = info[1]
		} else if info[0] == "hcl" {
			pass.hcl = info[1]
		} else if info[0] == "ecl" {
			pass.ecl = info[1]
		} else if info[0] == "pid" {
			pass.pid = info[1]
		} else if info[0] == "cid" {
			pass.cid = info[1]
		}
	}
	return pass
}
func getInput(fileName string) []Passport {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []Passport{}
	// Use Scan.
	p := ""
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 { // end of passport
			pass := parseFields(p)
			result = append(result, pass)
			p = ""
		} else {
			p += " " + line
		}
	}
	pass := parseFields(p)
	result = append(result, pass)
	return result
}
func part1(passports []Passport) {
	valid_passports := 0
	for _, passport := range passports {
		if passport.length_of_fields == 8 || (passport.length_of_fields == 7 && passport.cid == "") {
			valid_passports++
		}
	}
	fmt.Printf("Part1: %d\n", valid_passports)
}
func part2(passports []Passport) {
	valid_passports := 0
	for _, passport := range passports {
		reHeight := regexp.MustCompile("^(1[5-8][0-9]|19[0-3])cm$|^(59|6[0-9]|7[0-6])in$")
		reHairColor := regexp.MustCompile("^#[0-9a-f]{6}$")
		reEyeColor := regexp.MustCompile("^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$")
		rePassportId := regexp.MustCompile("^[0-9]{9}$")
		if passport.byr >= 1920 && passport.byr <= 2002 && passport.iyr >= 2010 && passport.iyr <= 2020 && passport.eyr >= 2020 && passport.eyr <= 2030 && reHeight.MatchString(passport.hgt) && reHairColor.MatchString(passport.hcl) && reEyeColor.MatchString(passport.ecl) && rePassportId.MatchString(passport.pid) {
			valid_passports++
		}
	}
	fmt.Printf("Part2: %d\n", valid_passports)
}
func main() {
	//part1(getInput("test_input.txt")) // test input
	part1(getInput("input.txt"))
	//part2(getInput("test_input2.txt")) // test input 2
	part2(getInput("input.txt"))
}

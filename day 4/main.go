package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var in []map[string]string
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passport := make(map[string]string)
	validcount := 0
	totalcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			in = append(in, passport)
			if checkPassport(passport) == 1 {
				if validatePassport(passport) {
					validcount++
				}
			}
			passport = make(map[string]string)
			totalcount++
			continue
		}
		values := strings.Split(line, ":")
		if _, ok := passport[values[0]]; ok {
			fmt.Printf("duplicate %v\n", values[0])
		}
		if values[0] == "cid" {
			continue
		}
		passport[values[0]] = values[1]
	}
	fmt.Printf("valid: %v\ntotal: %v\n", validcount, totalcount)
}

func checkPassport(pass map[string]string) int {
	if _, ok := pass["byr"]; !ok {
		return 0
	}
	if _, ok := pass["iyr"]; !ok {
		return 0
	}
	if _, ok := pass["eyr"]; !ok {
		return 0
	}
	if _, ok := pass["hgt"]; !ok {
		return 0
	}
	if _, ok := pass["hcl"]; !ok {
		return 0
	}
	if _, ok := pass["ecl"]; !ok {
		return 0
	}
	if _, ok := pass["pid"]; !ok {
		return 0
	}
	return 1
}

func validatePassport(passport map[string]string) bool {
	byr, err := strconv.Atoi(passport["byr"])
	if (err != nil) || byr < 1920 || byr > 2002 {
		//fmt.Printf("invalid byr: %v\n", passport["byr"])
		return false
	}
	iyr, err := strconv.Atoi(passport["iyr"])
	if (err != nil) || iyr < 2010 || iyr > 2020 {
		//fmt.Printf("invalid iyr: %v\n", passport["iyr"])
		return false
	}
	eyr, err := strconv.Atoi(passport["eyr"])
	if (err != nil) || eyr < 2020 || eyr > 2030 {
		//fmt.Printf("invalid eyr: %v\n", passport["eyr"])
		return false
	}
	height := passport["hgt"]
	hnumreg := regexp.MustCompile(`(\d+)(?:in|cm)`)
	hnum := hnumreg.FindStringSubmatch(height)
	if len(hnum) != 2 {
		return false
	}
	if strings.Contains(height, "in") {
		h := hnum[1]
		he, err := strconv.Atoi(h)
		if (err != nil) || he < 59 || he > 76 {
			return false
		}
	}
	if strings.Contains(height, "cm") {
		h := hnum[1]
		he, err := strconv.Atoi(h)
		if (err != nil) || he < 150 || he > 193 {
			return false
		}
	}
	creg := regexp.MustCompile(`(.*)(#[0-9a-z]{6})(.*)`)
	col := creg.FindStringSubmatch(passport["hcl"])
	if len(col) != 4 || (col[1] != "" || col[3] != "") {
		return false
	}
	ereg := regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	eye := ereg.FindStringSubmatch(passport["ecl"])
	if len(eye) != 2 {
		return false
	}
	preg := regexp.MustCompile(`(.*)([0-9]{9})(.*)`)
	pid := preg.FindStringSubmatch(passport["pid"])
	if len(pid) != 4 || (pid[1] != "" || pid[3] != "") {
		return false
	}
	return true
}

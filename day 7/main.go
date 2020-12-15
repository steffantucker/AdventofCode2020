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
	bags := map[string]string{}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	creg := regexp.MustCompile(`(.+) bags contain (.+)`)
	for scanner.Scan() {
		line := scanner.Text()
		info := creg.FindStringSubmatch(line)
		bags[info[1]] = info[2]
	}
	//foundBags := map[string]bool{}
	//findBag("shiny gold", bags, foundBags)
	//fmt.Printf("count: %v\n", len(foundBags))
	count := countBag("shiny gold", bags)
	fmt.Printf("bags reqd: %v\n", count)
}

func findBag(wanted string, b map[string]string, foundBags map[string]bool) {
	for bag, subbag := range b {
		if strings.Contains(subbag, wanted) {
			fmt.Printf("found %v in %v\n", wanted, bag)
			foundBags[bag] = true
			findBag(bag, b, foundBags)
		}
	}
}

func countBag(bag string, bags map[string]string) (count int) {
	if strings.Contains(bags[bag], "no other") {
		return 0
	}
	subbags := strings.Split(bags[bag], ", ")
	for _, subbag := range subbags {
		num, err := strconv.Atoi(string(subbag[0]))
		s := strings.Split(subbag, " bag")
		fmt.Printf("adding %v bags in %v\n", num, s[0][2:])
		if err != nil {
			panic(err)
		}
		count += num + num*countBag(s[0][2:], bags)
		fmt.Printf("count %v\n", count)
	}
	return
}

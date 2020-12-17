package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RulesFound struct {
	departure  bool
	field      int
	fieldfound bool
}

func main() {
	// part1()
	part2()
}

func part2() {
	rules := make(map[string]string)
	rulesfound := make(map[string]*RulesFound)
	tickets := make([][]string, 0)
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read instructions
	insreg := regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ins := insreg.FindStringSubmatch(line)
		rulesfound[ins[1]] = &RulesFound{departure: strings.Contains(ins[1], "departure")}
		findInclusive(ins[2], ins[3], ins[1], rules)
		findInclusive(ins[4], ins[5], ins[1], rules)
	}

	// our tickett
	scanner.Scan()
	ourticket := scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		errorrate := checkTicketString(strings.Split(line, ","), rules)
		if errorrate == 0 {
			n := strings.Split(line, ",")
			nums := []string{}
			for _, v := range n {
				nums = append(nums, v)
			}
			tickets = append(tickets, nums)
		}
	}
	found := findFields(tickets, rules, rulesfound)
	fmt.Println(found)
	departurenumber := findDep(found, strings.Split(ourticket, ","))
	fmt.Println(departurenumber)
}

func findDep(rules map[string]*RulesFound, ticket []string) int {
	dep := 1
	for _, v := range rules {
		if v.departure {
			num, _ := strconv.Atoi(ticket[v.field])
			dep *= num
		}
	}
	return dep
}

func findFields(tickets [][]string, rules map[string]string, rulesfound map[string]*RulesFound) map[string]*RulesFound {
	rcopy := make(map[string]string)
	revisit := []int{}
	for n, s := range rules {
		rcopy[n] = s
	}
	for j := 0; j < len(tickets[0]); j++ {
		r := make(map[string]int)
		for i := 0; i < len(tickets); i++ {
			for n, s := range rcopy {
				if strings.Contains(s, ","+tickets[i][j]+",") {
					r[n]++
				}
			}
		}
		biggest := ""
		bignum := 0
		for n, s := range r {
			if s == len(tickets) {
				bignum = s
				biggest = n
			}
			if bignum == len(tickets) && s == bignum {
				bignum = 0
				break
			}
		}
		if bignum < len(tickets) {
			revisit = append(revisit, j)
		} else {
			rulesfound[biggest].field = j
			rulesfound[biggest].fieldfound = true
			delete(rcopy, biggest)
		}
	}
	return rulesfound //revisitTix(tickets, revisit, rules, rulesfound)
}

//func revisitTix(tickets [])

func checkTicketString(ticket []string, rules map[string]string) (errorrate int) {
	str := ""
	for _, r := range rules {
		str += r
	}
	for _, v := range ticket {
		if !strings.Contains(str, v) {
			return 1
		}
	}
	return
}

func findInclusive(a, b, name string, rules map[string]string) {
	start, _ := strconv.Atoi(a)
	finish, _ := strconv.Atoi(b)
	str := ","
	for i := start; i <= finish; i++ {
		str += fmt.Sprintf("%v,", i)
	}
	rules[name] += str
}

/*func part1() {
	rules := make(map[int]string)
	rulesfound := make(map[string]RulesFound)
	tickets := make([][]int, 0)
	file, err := os.Open("inputtest")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read instructions
	insreg := regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ins := insreg.FindStringSubmatch(line)
		rulesfound[ins[1]] = RulesFound{departure: strings.Contains(ins[1], "departure")}
		findNums(ins[2], ins[3], ins[1], rules)
		findNums(ins[4], ins[5], ins[1], rules)
	}

	// our tickett
	scanner.Scan()
	scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		errorrate := checkTicket(strings.Split(line, ","), rules)
		if errorrate == 0 {
			n := strings.Split(line, ",")
			nums := []int{}
			for _, v := range n {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
			tickets = append(tickets, nums)
		}
	}
	found := findFields(tickets, rules, rulesfound)
	fmt.Println(found)
}*/

func checkTicket(ticket []string, rules map[int]string) (errorrate int) {
	for _, v := range ticket {
		a, _ := strconv.Atoi(v)
		if _, ok := rules[a]; !ok {
			errorrate += a
		}
	}
	return
}

func findNums(a, b, name string, rules map[int]string) {
	start, _ := strconv.Atoi(a)
	finish, _ := strconv.Atoi(b)
	for i := start; i <= finish; i++ {
		rules[i] = name
	}
}

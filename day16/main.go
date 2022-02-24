package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)

	conditions := []Condition{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		conditions = append(conditions, parseCond(scanner.Text()))
	}
	c := conditions[1]
	fmt.Println(c.From1, c.Name, c.To2)
	scanner.Scan()
	scanner.Scan()
	myTicket := scanner.Text()

	println("My ticket", myTicket)

	scanner.Scan()
	scanner.Scan()
	tickets := []string{}
	for scanner.Scan() {
		tickets = append(tickets, scanner.Text())
	}

	sum := 0
	for i, ticket := range tickets {
		tsum := incorrect(ticket, conditions)
		println(i, tsum)
		sum += tsum
	}

	println("Sum", sum)
}

func incorrect(ticket string, conds []Condition) int {
	spl := strings.Split(ticket, ",")

	nums := make([]int, len(spl))
	for i, s := range spl {
		nums[i], _ = strconv.Atoi(s)
	}

	sum := 0
	for _, num := range nums {
		valid := false
		for _, cond := range conds {
			if cond.Valid(num) {
				valid = true
				break
			}
		}
		println("valid", valid,  num)
		if !valid {
			sum += num
		}
	}
	return sum
}


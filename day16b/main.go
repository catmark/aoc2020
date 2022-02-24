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
	scanner.Scan()
	scanner.Scan()
	myTicketStr := strings.Split(scanner.Text(), ",")	
	myTicket := make([]int, len(myTicketStr))
	for i, s := range myTicketStr {
		myTicket[i], _ = strconv.Atoi(s)
	}

	println("My ticket", myTicket)

	scanner.Scan()
	scanner.Scan()
	tickets := []string{}
	for scanner.Scan() {
		tickets = append(tickets, scanner.Text())
	}

	validTickets := [][]int{}
	for _, ticket := range tickets {
		vt, valid := incorrect(ticket, conditions)
		if valid {
			validTickets = append(validTickets, vt)
		}
	}
	validTickets = append(validTickets, myTicket)

	println("Valid tickets", len(validTickets))

	possibilities := mapFields(validTickets, conditions)
	for i, field := range possibilities {
		fmt.Println(i, ":", len(field), "options:", strings.Join(field, ","))
	}

	final := make([]string, 20)
	taken := []string{}
	for len(taken) < 20 {
		for i, options := range possibilities {
			available := []string{}
			for _, option := range options {
				if !inLst(option, taken) {
					available = append(available, option)
				}
			}
			if len(available) == 1 {
				taken = append(taken, available[0])
				final[i] = available[0]
				fmt.Println("Setting", available[0], "at", i)
			}
		}
	}
	total := 1
	for i, t := range final {
		fmt.Println(i, t)
		if strings.HasPrefix(t, "departure") {
			total *= myTicket[i]
		}
	}
	fmt.Println("Total", total)
}

func inLst(s string, lst []string) bool {
	for _, l := range lst {
		if l == s {
			return true
		}
	}
	return false
}

func incorrect(ticket string, conds []Condition) ([]int, bool) {
	spl := strings.Split(ticket, ",")

	nums := make([]int, len(spl))
	for i, s := range spl {
		nums[i], _ = strconv.Atoi(s)
	}

	for _, num := range nums {
		valid := false
		for _, cond := range conds {
			if cond.Valid(num) {
				valid = true
				break
			}
		}
		if !valid {
			return nums, false
		}
	}
	return nums, true
}

func mapFields(tickets [][]int, conditions []Condition) [][]string {
	fields := make([][]int, 20) // group values by field instead of by ticket

	for _, t := range tickets {
		for i, num := range t {
			fields[i] = append(fields[i], num)
		}
	}

	sorted := make([][]string, 20)

	for i, field := range fields {
		for _, cond := range conditions {
			allValid := true
			for _, num := range field {
				if !cond.Valid(num) {
					allValid = false
				}
			}
			if allValid {
				sorted[i] = append(sorted[i], cond.Name)
			}
		}
	}
	return sorted
}


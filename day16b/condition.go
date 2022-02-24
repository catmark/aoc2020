package main

import (
	"strconv"
	"strings"
)

type Condition struct {
	Name string
	From1 int
	To1 int
	From2 int
	To2 int
}

func (c Condition) Valid(i int) bool {
	return (i >= c.From1 && i <= c.To1) || (i >= c.From2 && i <= c.To2) 
}

func parseCond(input string) Condition {
	c := Condition{}

	spl := strings.Split(input, ":")
	c.Name = spl[0]
	sets := strings.Split(spl[1], "or")
	c.From1, c.To1 = fromTo(sets[0])
	c.From2, c.To2 = fromTo(sets[1])
	return c

}

func fromTo(s string) (int, int) {
	spl := strings.Split(strings.TrimSpace(s), "-")

	from, _ := strconv.Atoi(spl[0])
	to, _ := strconv.Atoi(spl[1])

	return from, to
}


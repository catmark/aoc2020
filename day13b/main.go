package main

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
	"strings"
)



func main() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	busString := strings.Split(scanner.Text(), ",")
	busses := map[int]int64{}
	for i, s := range busString {
		if s != "x" {
			bus, _ := strconv.Atoi(s)
			busses[i] = int64(bus)
		}
	}
	fmt.Println("Busses", busses)

	fmt.Println(firstSync2(busses))

}

func firstBus(timestamp int, busses []int) (int, int) {
	for {
		for _, bus := range busses {
			if timestamp % bus == 0 {
				return bus, timestamp
			}
		}
		timestamp++
	}
	return 0, 0
}

func firstSync(busses map[int]int64) int64 {
	for ts := int64(1000); ; ts += 983 {
		if ts % 1000000000 == 0 {
			println(ts)
		}
		ok := true
		for i, bus := range busses {
			if (ts + int64(i)) % bus != 0 {
				ok = false
				break
			}
		}
		if ok {
			return ts
		}
	}
}

func firstSync2(busses map[int]int64) int64 {
	stepSize := int64(1)
	timestamp := int64(1)

	for i, bus := range busses {
		for ;; timestamp += stepSize {
			if (timestamp + int64(i)) % bus == 0 {
				fmt.Println("Matched bus", bus, "at", timestamp)
				break;
			}
		}
		stepSize *= bus
		fmt.Println("Setting stepsize to", stepSize)
	}
	return timestamp
}

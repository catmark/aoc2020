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

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	memory := run(lines)

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

func run(lines []string) map[int64]int64 {
	mask := []byte{}
	memory := map[int64]int64{}

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = parseMask(line)
		} else {
			spl := strings.Split(line, "] = ")
			mem := strings.Split(spl[0], "[")
			address, _ := strconv.ParseInt(mem[1], 10, 64)
			value, _ := strconv.ParseInt(spl[1], 10, 64)

			addresses := applyMask(address, mask)
			for _, a := range addresses {
				fmt.Println("Setting", a, "to", value)
				memory[a] = value
			}
		}
	}
	return memory
}

func parseMask(mask string) []byte {
	spl := strings.Split(mask, " = ")
	return []byte(spl[1])
}


func applyMask(address int64, mask []byte) []int64 {
	n := strconv.FormatInt(address, 2)
	num := make([]byte, 36)

	numLen := len(n)
	for i := 0; i < 36; i++ {
		if i < 36 - numLen {
			num[i] = '0'
		} else {
			num[i] = n[i - (36 - numLen)]
		}
	}
	for i, bit := range mask {
		if bit != '0' {
			num[i] = bit
		}
	}
	return makeAddresses(string(num))
}

func makeAddresses(masked string) []int64 {
	index := strings.IndexByte(masked, 'X')

	if index == -1 {
		x, _ := strconv.ParseInt(string(masked), 2, 64)
		return []int64{x}
	}

	copy0 := []byte(masked)
	copy0[index] = '0'
	with0 := makeAddresses(string(copy0))
	copy1 := []byte(masked)
	copy1[index] = '1'
	with1 := makeAddresses(string(copy1))
	return append(with0, with1...)
}

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

func run(lines []string) map[int]int64 {
	mask := []byte{}
	memory := map[int]int64{}

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = parseMask(line)
		} else {
			spl := strings.Split(line, "] = ")
			mem := strings.Split(spl[0], "[")
			address, _ := strconv.Atoi(mem[1])
			value, _ := strconv.ParseInt(spl[1], 10, 64)

			memory[address] = applyMask(value, mask)

		}
	}
	return memory


}

func parseMask(mask string) []byte {
	spl := strings.Split(mask, " = ")
	return []byte(spl[1])
}


func applyMask(value int64, mask []byte) int64 {
	n := strconv.FormatInt(value, 2)
	println(n)
	num := make([]byte, 36)
	numLen := len(n)
	for i := 0; i < 36; i++ {
		if i < 36 - numLen {
			num[i] = '0'
		} else {
			num[i] = n[i - (36 - numLen)]
		}
	}
	println(string(num), "<=")
	for i, bit := range mask {
		if bit != 'X' {
			num[i] = bit
		}
	}

	x, _ := strconv.ParseInt(string(num), 2, 64) 
	return x
}




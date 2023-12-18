package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)
	total := 0

	for scan.Scan() {
		line := scan.Text()

		first := 0
		last := 0

		build := ""

		for _, char := range line {
			found := 0
			if unicode.IsDigit(char) {
				found = int(char - '0')
				build = ""
			} else {
				build += string(char)

				if strings.Contains(build, "one") {
					found = 1
					build = "e"
				} else if strings.Contains(build, "two") {
					found = 2
					build = "o"
				} else if strings.Contains(build, "three") {
					found = 3
					build = "e"
				} else if strings.Contains(build, "four") {
					found = 4
					build = ""
				} else if strings.Contains(build, "five") {
					found = 5
					build = "e"
				} else if strings.Contains(build, "six") {
					found = 6
					build = ""
				} else if strings.Contains(build, "seven") {
					found = 7
					build = ""
				} else if strings.Contains(build, "eight") {
					found = 8
					build = "t"
				} else if strings.Contains(build, "nine") {
					found = 9
					build = "e"
				}
			}
			if found != 0 {
				// fmt.Printf("%v -- %v\n", string(char), build)
				if first == 0 {
					first = found
				}
				last = found
			}
		}
		total += first*10 + last
		// fmt.Printf("line: %v -- first: %v -- last: %v\n", line, first, last)
		// fmt.Printf("value: %v -- total: %v\n", first*10+last, total)
	}

	fmt.Println(total)

}

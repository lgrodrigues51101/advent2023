package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BagContent struct {
	red   int
	green int
	blue  int
}

func main() {

	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	// content := BagContent{
	// 	red:   12,
	// 	green: 13,
	// 	blue:  14,
	// }

	totalpower := 0

	for scan.Scan() {
		line := scan.Text()

		minred, mingreen, minblue := 1, 1, 1

		split1 := strings.Split(line, ":")

		// id, _ := strconv.Atoi(strings.Split(split1[0], " ")[1])

		split2 := strings.Split(split1[1], ";")

		possible := true
		// outerLoop:
		for _, set := range split2 {

			split3 := strings.Split(set, ",")

			for _, move := range split3 {

				split4 := strings.Split(move, " ")

				value, _ := strconv.Atoi(split4[1])
				color := split4[2]

				// // test if the move is possible
				// if (color == "red" && value > content.red) ||
				// 	(color == "green" && value > content.green) ||
				// 	(color == "blue" && value > content.blue) {
				// 	possible = false
				// 	break outerLoop
				// }

				// evaluate if the minimun should be updated
				if color == "red" && value > minred {
					minred = value
				} else if color == "green" && value > mingreen {
					mingreen = value
				} else if color == "blue" && value > minblue {
					minblue = value
				}
			}
		}
		if possible {
			// fmt.Printf("%v-%v-%v-%v\n", id, minred, mingreen, minblue)
			totalpower += minred * mingreen * minblue
		}
	}
	fmt.Println(totalpower)
}

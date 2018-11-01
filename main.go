package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type city struct {
	name string
	temp int
}

func main() {
	var noCities int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		noCities, err = strconv.Atoi(scanText(scanner, "❔  Number of cities: "))
		if err == nil {
			break
		}
		fmt.Println("❗️  Invalid number of cities")
	}

	highestTempIdx := 0
	c := make([]city, noCities)
	for i := 0; i < noCities; i++ {
		c[i].name, c[i].temp = scanNameAndTemp(scanner, i+1)
		if i != 0 {
			if c[i].temp > c[highestTempIdx].temp {
				highestTempIdx = i
			}
		}
	}

	fmt.Printf("☀️   Highest temperature is %d in %s\n\n", c[highestTempIdx].temp, c[highestTempIdx].name)
}

func scanText(scanner *bufio.Scanner, question string) string {
	var scan string
	for {
		fmt.Print(question)
		if scanner.Scan() {
			scan = scanner.Text()
		}
		if err := scanner.Err(); err == nil {
			break
		}
		// ERROR
	}
	return scan
}

func scanNameAndTemp(scanner *bufio.Scanner, idx int) (name string, temp int) {
	var nameScan, tempScan string
	var err error

	for {
		nameQst := fmt.Sprintf("🏙   Name of city #%d: ", idx)
		nameScan = scanText(scanner, nameQst)
		name = strings.TrimSpace(nameScan)
		if len(name) > 0 {
			break
		}
		fmt.Println("❗️  Invalid city name")
	}

	for {
		tempQst := fmt.Sprintf("🌡   Temperature in %s: ", nameScan)
		tempScan = scanText(scanner, tempQst)
		temp, err = strconv.Atoi(strings.TrimSpace(tempScan))
		if err == nil {
			break
		}

		fmt.Println("❗️  Invalid temperature")
	}

	return name, temp
}

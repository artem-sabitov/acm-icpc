package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const plusOne = "+one"
const superstitiousNumber = 13
const costPerPerson = 100

func main() {
	var totalPerson int64 = 2 // by default 2 guest

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	guestCount, error := strconv.ParseInt(scanner.Text(), 10, 64)
	if error != nil {
		fmt.Println(error)
	}

	for i := int64(0); i < guestCount; i++ {
		scanner.Scan()
		invite := scanner.Text()
		if strings.Contains(invite, plusOne) {
			totalPerson += 2
		} else {
			totalPerson++
		}
	}

	if totalPerson == superstitiousNumber {
		totalPerson++
	}

	dinnerCost := totalPerson * costPerPerson

	fmt.Println(dinnerCost)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numbers int64
var primeNumberDepth int64

const maxNumbers int64 = 16256

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numbers, error := strconv.ParseInt(scanner.Text(), 10, 64)
	if error != nil {
		fmt.Println(error)
		return
	}

	var primeNumberPositions [maxNumbers]int64

	for i := int64(0); i < numbers; i++ {
		scanner.Scan()
		currentPosition, error := strconv.ParseInt(scanner.Text(), 10, 64)
		if error != nil {
			fmt.Println(error)
			return
		}

		primeNumberPositions[i] = currentPosition

		if currentPosition > primeNumberDepth {
			primeNumberDepth = currentPosition
		}
	}

	primeNumbers := calcPrimeNumberSequence(maxNumbers)
	for i := int64(0); i < maxNumbers; i++ {
		currentPosition := primeNumberPositions[i]
		if currentPosition == 0 {
			break
		}

		fmt.Println(primeNumbers[currentPosition-1])
	}
}

func calcPrimeNumberSequence(maxNumber int64) [maxNumbers]int64 {
	var foundNumbers [maxNumbers]int64
	totalFound := 0

	for i := int64(2); i <= maxNumber; i++ {
		isPrime := true
		for j := int64(2); j <= maxNumber; j++ {
			if i != j {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
		}

		if isPrime {
			index := totalFound
			foundNumbers[index] = i
			totalFound++
		}
	}

	// for i := 0; i < totalFound; i++ {
	// 	fmt.Println(foundNumbers[i])
	// }

	return foundNumbers
}

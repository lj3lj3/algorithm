package jz

import "fmt"

func Print1ToMaxOfNDigits(n int) {
	if n < 0 {
		return
	}

	// init n length long byte array
	digits := make([]byte, n, n)
	loop := true
	for loop {
		// print number
		printDigits(digits)

		// increase lowest digit
		digits[len(digits)-1]++

		// increase next if needed
		var increaseNext bool
		for i := len(digits) - 1; i >= 0; i-- {
			// increase next
			if increaseNext {
				digits[i]++
				increaseNext = false
			}
			if digits[i] > 9 {
				if i == 0 {
					loop = false
					break
				}
				// need increase number
				increaseNext = true
				digits[i] -= 10
			}
		}
	}

	fmt.Printf("\n")
}

func printDigits(digits []byte) {
	var printZero bool
	var printAny bool
	for _, digit := range digits {
		if digit == 0 {
			if !printZero {
				continue
			}
		} else if !printZero {
			printZero = true
		}

		printAny = true
		fmt.Printf("%d", digit)
	}

	// print divider
	if printAny {
		fmt.Printf(", ")
	}
}

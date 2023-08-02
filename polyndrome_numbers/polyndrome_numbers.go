package polyndrome_numbers

import "fmt"

func FindReverseNumber(n uint64) uint64 {
	var acc uint64
	var accCount uint64

	var i uint64
	for {
		rn := isRevNumber(i)

		if rn {
			acc = i
			accCount++
			fmt.Printf("i, acc=%d; accCount=%d\n", acc, accCount)
		}
		if accCount == n {
			break
		}
		i++
	}
	return acc
}

func isRevNumber(n uint64) bool {
	if n < 10 {
		return true
	}
	numbers := make([]uint64, 0)
	for {
		a := n % 10
		numbers = append(numbers, a)
		n = n / 10
		if n < 10 {
			numbers = append(numbers, n)
			break
		}
	}
	for i := 0; i < len(numbers)/2; i++ {
		fd := numbers[i]
		ld := numbers[len(numbers)-1-i]
		if fd != ld {
			return false
		}
	}
	return true
}

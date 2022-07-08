package main

import "fmt"

type evenodd []string

func zeroToN(n int) []int {

	numbers := []int{}

	for i := 0; i < n+1; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func evenOrOdd(n []int) evenodd {

	output := evenodd{}

	for _, num := range n {
		if num%2 != 0 {
			output = append(output, "is odd")
		} else {
			output = append(output, "is even")
		}
	}

	return output
}

func (e evenodd) print() {

	for i, num := range e {
		fmt.Println(i, num)
	}
}

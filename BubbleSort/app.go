package main

import "fmt"

func sweep(numbers []int, prevPasses int)  {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < (N-prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber{
			// Swap the numbers!
			numbers[firstIndex] = secondNumber
			numbers[secondIndex]=firstNumber
		}
		firstIndex++
		secondIndex++
	}
}

func bubbleSort(numbers []int)  {
	var N int = len(numbers)
	var i int
	for i=0;i<N;i++{
		sweep(numbers, i)
	}
}

func main() {
	var numbers []int = []int{5,4,3,1,2}
	fmt.Println("Out list of numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
}

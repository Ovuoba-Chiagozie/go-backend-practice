package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumALL(numbersToPass ...[]int) []int {

	// lengthOfNumbers := len(numbersToPass)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToPass {
	// 	sums[i] = Sum(numbers)
	// }

	// return sums

	var sums []int

	for _, numbers := range numbersToPass {

		sums = append(sums, Sum(numbers))
	}

	return sums
}

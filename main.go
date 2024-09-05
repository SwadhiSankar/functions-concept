package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) // no parenthesis for double -> to passing func and not returning the result of the func if we give () it will return the result
	triple := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(triple)
}

func transformNumbers(numbers *[]int, tranform func(int)int)[]int{
	dNumber := []int{}

	for _, value := range *numbers{
		dNumber = append(dNumber,tranform(value))
	}

	return dNumber

	
}
func double(number int ) int{
	return number * 2
}

func triple(number int) int{
	return number * 3
}
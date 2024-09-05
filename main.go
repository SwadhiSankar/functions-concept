package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumber(&numbers)

	fmt.Println(doubled)
}

func doubleNumber(numbers *[]int)[]int{
	dNumber := []int{}

	for _, value := range *numbers{
		dNumber = append(dNumber,value * 2)
	}

	return dNumber

}
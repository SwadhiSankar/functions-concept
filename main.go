package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := [] int {5,1,2,3}
	doubled := transformNumbers(&numbers, double) // no parenthesis for double -> to passing func and not returning the result of the func if we give () it will return the result
	triple := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(triple)

	transformFn1 := getTransformerFunction(&numbers) //2
	transformFn2 := getTransformerFunction(&moreNumbers) //15

    transfomredNumbers := transformNumbers(&numbers, transformFn1)
    moreTransformedNumbers := transformNumbers(&numbers, transformFn2)

	fmt.Print(transfomredNumbers)
	fmt.Print(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, tranform func(int)int)[]int{
	dNumber := []int{}

	for _, value := range *numbers{
		dNumber = append(dNumber,tranform(value))
	}

	return dNumber

	
}

func getTransformerFunction(numbers *[] int)func(int)int {
    if(*numbers)[0] ==1{
		return double
	}else{
	return triple 
	}// returning a method not the result
}
func double(number int ) int{
	return number * 2
}

func triple(number int) int{
	return number * 3
}
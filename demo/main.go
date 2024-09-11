package main

import "fmt"

type myFunc func(int, int) int

type Flight struct {
	number      int
	airlineCode string
	departure   string
}

func main() {
	sumResult := calculate(sum, 3, 4)
	fmt.Println(sumResult)

	minusResult := calculate(minus, 3, 4)
	fmt.Println(minusResult)

	flight := Flight{}
	fmt.Printf("%#v\n", flight)

	myFlight := Flight{number: 45, airlineCode: "CB", departure: "13.00"}
	fmt.Printf("%#v\n", myFlight)
	fmt.Printf("number: %v", myFlight.number)
}

func calculate(fn myFunc, a, b int) int {
	return fn(a, b)
}

func sum(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

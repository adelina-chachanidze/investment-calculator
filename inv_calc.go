package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	const inflationRate = 2.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment years: ")
	fmt.Scan(&years)

	fmt.Print("Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value:", futureValue)
	fmt.Println("Future value with inflation:", futureRealValue)
}

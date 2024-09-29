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

	// println method
	//fmt.Println("Future value:", futureValue)
	//fmt.Println("Future value with inflation:", futureRealValue)

	//Printf method
	//fmt.Printf("Future value: %.1f \nFuture value with inflation: %.1f", futureValue, futureRealValue)

	//Print with vars and sprintf method
	formattedFV := fmt.Sprintf("Future value: %.1f \n", futureValue)
	formattedRV := fmt.Sprintf("Future value with inflation: %.1f \n", futureRealValue)
	fmt.Print(formattedFV, formattedRV)
}

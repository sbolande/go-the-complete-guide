package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount := promptUserForFloat("Investment Amount")
  expectedReturnRate := promptUserForFloat("Expected Return Rate")
	years := promptUserForFloat("Years")

	futureValue, futureRealValue := calcFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: $%.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): $%.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}

func promptUserForFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt, ": ")
	fmt.Scan(&input)
	return input
}
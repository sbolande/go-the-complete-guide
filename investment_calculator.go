package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: $%.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): $%.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}
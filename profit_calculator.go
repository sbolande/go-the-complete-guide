package main

import "fmt"

func main() {
	revenue := getUserInputFloat("Revenue")
	expenses := getUserInputFloat("Expenses")
	taxRate := getUserInputFloat("Tax Rate")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInputFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt, ": ")
	fmt.Scan(&input)
	return input
}
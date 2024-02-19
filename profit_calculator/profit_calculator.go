package profit_calculator

import (
	"errors"
	"fmt"
	"os"
)

const financialsFileName = "finances.txt"

func main() {
	revenue, err1 := getUserInputFloat("Revenue")
	expenses, err2 := getUserInputFloat("Expenses")
	taxRate, err3 := getUserInputFloat("Tax Rate")
	if err1 != nil || err2 != nil || err3 != nil { 
		fmt.Println(err1, err2, err3)
		return
	}
	
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeFinancialsToFile(ebt, profit, ratio)
}

func getUserInputFloat(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt, ": ")
	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("Value must be a positive number")
	}
	return input, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func writeFinancialsToFile(ebt, profit, ratio float64) {
	financialsText := fmt.Sprintf("EBT: %.1f | PROFIT: %.1f | RATIO: %.3f", ebt, profit, ratio)
	fmt.Println(financialsText)
	os.WriteFile(financialsFileName, []byte(financialsText), 0644)
}
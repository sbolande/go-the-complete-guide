package bank

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFileName)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}
  balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance data")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmt float64
			fmt.Scan(&depositAmt)
			if depositAmt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmt
			fmt.Println("Balance updated. New balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmt float64
			fmt.Scan(&withdrawalAmt)
			if withdrawalAmt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			} else if withdrawalAmt > accountBalance {
				fmt.Println("Invalid amount. You may not withdraw more than is in your account.")
				continue
			}
			accountBalance -= withdrawalAmt
			fmt.Println("Balance updated. New balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodybe!")
			fmt.Println("Thanks for choosing Go Bank!")
			return
		}
	}
}
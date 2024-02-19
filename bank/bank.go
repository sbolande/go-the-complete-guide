package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFileName = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
	
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
		default:
			fmt.Println("Goodybe!")
			fmt.Println("Thanks for choosing Go Bank!")
			return
		}
	}
}
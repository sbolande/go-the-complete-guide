package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
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
		default:
			fmt.Println("Goodybe!")
			fmt.Println("Thanks for choosing Go Bank!")
			return
		}
	}
}
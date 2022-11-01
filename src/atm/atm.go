package atm

import (
	"fmt"
	"os"
	"strconv"
	// "strconv"
)

var (
	pin        = 1234
	defaultPin = 1234
	balance    = 10000
	amount     int
	choice     int
)

func StartBanking() {
	fmt.Println("Welcome to the ATM")
	fmt.Println("Please enter your pin to continue (Default pin for new users : 1234")
	var pinInput int
	fmt.Scanln(&pinInput)
	if pinInput == defaultPin {
		fmt.Println("Welcome")
		fmt.Println("Please Change your default Pin to continue")
		changePin()
	} else if pinInput == pin && len(strconv.Itoa(pinInput)) == 4 {
		fmt.Println("Welcome")
		startBanking()
	} else {
		fmt.Println("Invalid Pin")
		fmt.Println("Please try again")
	}
}

// create function to change pin

func changePin() {
	fmt.Println("Enter your new pin")
	var newPin int
	fmt.Scanln(&newPin)
	if newPin == defaultPin {
		fmt.Println("You can't use the default pin")
		changePin()
	} else if len(strconv.Itoa(newPin)) == 4 {
		pin = newPin
		fmt.Println("Your pin has been changed successfully!")
		fmt.Println("")
		startBanking()
	} else {
		fmt.Println("Pin should be 4 digits")
		changePin()
	}
}

// create function to check account balance
func checkBalance() {
	fmt.Println("Your account balance is", strconv.Itoa(balance))
	continueBanking()
}

// create function to withdraw funds
func withdrawFunds() {
	fmt.Println("How much would you like to withdraw?")

	fmt.Scanln(&amount)
	if amount > balance {
		fmt.Println("Insufficient funds")
	} else if amount < 0 {
		fmt.Println("Invalid amount")
	} else {
		balance = balance - amount
		fmt.Println("Withdrawal successful")
		fmt.Println("Your new balance is", strconv.Itoa(balance))
	}
	continueBanking()
}

// create function to deposit funds
func depositFunds() {
	fmt.Println("How much would you like to deposit?")
	fmt.Scanln(&amount)
	if amount > 0 {
		balance = balance + amount
		fmt.Println("Deposit successful")
		fmt.Println("Your new balance is", strconv.Itoa(balance))
	} else {
		fmt.Println("Invalid amount")
	}
	continueBanking()
}

// create function to exit
func exit() {
	fmt.Println("Thank you for banking with us")
	os.Exit(0)
}

// create a function to display menu
func displayMenu() {
	fmt.Println("Please select an option")
	fmt.Println("")
	fmt.Println("1. Change Pin")
	fmt.Println("2. Account Balance")
	fmt.Println("3. Withdraw Funds")
	fmt.Println("4. Deposit Funds")
	fmt.Println("5. Cancel/Exit")

}

// create a function to perform selected operation
func performOperation(choice int) {
	switch choice {
	case 1:
		changePin()
	case 2:
		checkBalance()
	case 3:
		withdrawFunds()
	case 4:
		depositFunds()
	case 5:
		exit()
	default:
		fmt.Println("Invalid option selected")

	}
}

// create a function to start banking
func startBanking() {

	displayMenu()
	fmt.Scanln(&choice)
	performOperation(choice)

}

func continueBanking() {
	fmt.Println("")
	fmt.Println("Would you like to perform another operation?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Scanln(&choice)
	if choice == 1 {
		startBanking()
	} else {
		exit()
	}
}

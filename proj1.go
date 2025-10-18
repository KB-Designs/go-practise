package main

import "fmt"

/*
Calculate Balance

We need to calculate the total cost for a batch of messages, and update the user’s balance if they have enough money.
Assignment

Using the given variables, write conditional statements to calculate and update the variables.

    Set finalCost to the bulkMessageCost.
    If the user is a premium user, apply the discountRate to the finalCost.
        For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.
    If the user has enough money in their accountBalance:
        Deduct finalCost from their accountBalance.
        Print the purchaseSuccessMessage
    If not, just print the insufficientFundMessage.

Assignment

Using the given variables, write conditional statements to calculate and update the variables.

    Set finalCost to the bulkMessageCost.
    If the user is a premium user, apply the discountRate to the finalCost.
        For example, a discountRate of 0.10 means the discounted price per message would be 90% of the original price.
    If the user has enough money in their accountBalance:
        Deduct finalCost from their accountBalance.
        Print the purchaseSuccessMessage
    If not, just print the insufficientFundMessage.


*/



func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line
	
	if isPremiumUser {
		finalCost=bulkMessageCost-(bulkMessageCost*discountRate)
	} else {
		finalCost=bulkMessageCost
	}

	if finalCost<= accountBalance {
		accountBalance=accountBalance-finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}


	fmt.Println("Account balance:", accountBalance)
}

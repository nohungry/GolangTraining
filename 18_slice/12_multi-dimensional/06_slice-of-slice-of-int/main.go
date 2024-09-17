package main

import (
	"fmt"
)

func main() {

	transactions := make([][]int, 0, 3)
	fmt.Println(transactions)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		fmt.Println(transaction)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

// Name: Syed Zulkifal Banuri
// RollNo: 20I-2350
// Assignment 2

package main

import (
	"fmt"

	"github.com/syed199921/assignment02bca"
)

func main() {
	Blockchain := assignment02bca.BlockChain{}
	var minBlockHash string = ""
	var maxBlockHash string = ""
	var mainOption int = -1
	var numberOfTransactionsPerBlock int
	fmt.Println("Welcome to Syed's Blockchain")
	fmt.Println("-----------------------------")
	//Hash range for the values:
	//1. 0000000000000000000000000000000000000000000000000000000000000000 (the minimum value in hexadecimals)
	//2. FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF (the maximum value in hexadecimals)
	// Blockchain.SetBlockHashRange("66687AADF862BD776C8FC18B8E9F8E20089714856EE233B3902A591D0D5F2925", "AF9613760F72635FBDB44A5A0A63C39F12AF30F950A6EE5C971BE188E89C4051")
	// Blockchain.SetNumberOfTransactionsPerBlock(2)
	fmt.Println("Enter the minimum value of the hash range:")
	fmt.Scanln(&minBlockHash)
	fmt.Println("Enter the maximum value of the hash range:")
	fmt.Scanln(&maxBlockHash)
	Blockchain.SetBlockHashRange(minBlockHash, maxBlockHash)
	for {
		fmt.Println("Enter the number of transactions per block:")
		fmt.Scanln(&numberOfTransactionsPerBlock)
		if numberOfTransactionsPerBlock < 0 {
			fmt.Println("Please enter a positive number")

		} else if numberOfTransactionsPerBlock == 0 {
			fmt.Println("Please enter a number greater than 0")
		} else {
			break
		}
	}
	Blockchain.SetNumberOfTransactionsPerBlock(numberOfTransactionsPerBlock)

	for {
		fmt.Println("Main Menu\n------------------------\nPlease select an option:\n1. Add a transaction\n2. Display the blockchain\n3. Change a transaction in a block\n4. Verify blockchain\n5. Exit Blockchain")
		fmt.Scanln(&mainOption)
		switch mainOption {
		case 1:
			{
				var transaction string
				fmt.Println("Enter the transaction")
				fmt.Scanln(&transaction)
				Blockchain.AddTransaction(transaction)
				var option int
				for {
					fmt.Println("Enter 0 to return to main menu")
					fmt.Scanln(&option)
					if option == 0 {
						break
					}
				}
			}

		case 2:
			{
				Blockchain.DisplayBlockchain()
				var option int
				for {
					fmt.Println("Enter 0 to return to main menu")
					fmt.Scanln(&option)
					if option == 0 {
						break
					}
				}
			}
		case 3:
			{
				var blockHash, transactionHash, newTransaction string = "", "", ""
				fmt.Println("Enter the hash of the block to change:")
				fmt.Scanln(&blockHash)
				fmt.Println("Enter the hash of the transaction to change:")
				fmt.Scanln(&transactionHash)
				fmt.Println("Enter the new transaction")
				fmt.Scanln(&newTransaction)
				Blockchain.ChangeBlock(transactionHash, newTransaction, blockHash)
				var option int
				for {
					fmt.Println("Enter 0 to return to main menu")
					fmt.Scanln(&option)
					if option == 0 {
						break
					}
				}
			}
		case 4:
			{
				if Blockchain.VerifyBlockchain() == true {
					fmt.Println("Blockchain is valid")
				} else {
					fmt.Println("Therefore the blockchain is not valid")
				}
				var option int
				for {
					fmt.Println("Enter 0 to return to main menu")
					fmt.Scanln(&option)
					if option == 0 {
						break
					}
				}
			}

		}

		if mainOption == 5 {
			fmt.Println("See you soon! Goodbye!")
			break
		} else {
			continue
		}

	}

}

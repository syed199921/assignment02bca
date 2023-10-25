package main

import (
	"github.com/syed199921/assignment02bca"
)

func main() {
	Blockchain := assignment02bca.BlockChain{}
	transactions1 := []string{"2 btc to syed", "1 btc to hamza", "3 btc to syed"}
	transactions2 := []string{"2 btc to syed2", "1 btc to hamza2", "3 btc to syed2"}
	transactions3 := []string{"2 btc to syed3", "1 btc to hamza4", "3 btc to syed3"}
	Blockchain.AddBlock(transactions1, 1)
	Blockchain.AddBlock(transactions2, 2)
	Blockchain.AddBlock(transactions3, 3)
	Blockchain.DisplayBlockchain()

	// fmt.Println(Blockchain.VerifyTransaction("f57678f87f5bbe43a60190e74b87dee999da07de2d57b8f9818b1fdbfead2709", "d75461dd72c6c0a0fdd3f4a4ecfa3773e1a6301f7eca9ff3f078291c95372969"))
}

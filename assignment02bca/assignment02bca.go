package assignment02bca

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	hash  string
}

var transactions []string = nil

type Block struct {
	transactions      []string
	nonce             int
	previousHash      string
	hash              string
	timeOfTransaction string
	previous          *Block
	merkleRoot        *MerkleNode
}

type BlockChain struct {
	tail                         *Block
	numberOfTransactionsPerBlock int
	minBlockHash                 string
	maxBlockHash                 string
}

func CalculateHash(stringToHash string) string {
	h := sha256.New()
	h.Write([]byte(stringToHash))
	return string(h.Sum(nil))
}

func (b *Block) buildMerkleTree() *MerkleNode {
	var nodes []*MerkleNode

	for _, transaction := range b.transactions {
		nodes = append(nodes, &MerkleNode{nil, nil, CalculateHash(transaction)})
	}

	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		var newNodes []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			hash := CalculateHash(fmt.Sprintf("%x", []byte(left.hash)) + fmt.Sprintf("%x", []byte(right.hash)))
			newNode := &MerkleNode{left, right, hash}
			newNodes = append(newNodes, newNode)

		}

		nodes = newNodes
	}

	return nodes[0]

}
func generateRandomNonce() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}

func (bc *BlockChain) AddTransaction(transaction string) {
	transactions = append(transactions, transaction)
	fmt.Println("Transaction added successfully")
	fmt.Println("Number of transactions in the block: ", len(transactions))
	if len(transactions) == bc.numberOfTransactionsPerBlock {
		bc.AddBlock(transactions, generateRandomNonce())
		transactions = nil
	}
}

func (bc *BlockChain) AddBlock(transactions []string, nonce int) {

	newBlock := &Block{transactions, nonce, "", "", time.Now().String(), bc.tail, nil}
	if bc.tail != nil {
		newBlock.previousHash = bc.tail.hash
	} else {

		newBlock.previousHash = ""
	}
	newBlock.merkleRoot = newBlock.buildMerkleTree()
	for {
		str := newBlock.merkleRoot.hash + strconv.Itoa(newBlock.nonce) + newBlock.previousHash + newBlock.timeOfTransaction
		hash := CalculateHash(str)

		if hash >= bc.minBlockHash && hash <= bc.maxBlockHash {
			newBlock.hash = hash
			break
		} else {
			fmt.Println("hash not in range, generating new nonce and trying again")
			newBlock.nonce = generateRandomNonce()
		}
	}
	bc.tail = newBlock
}

func displayMerkleTree(node *MerkleNode, depthOfTree int, isLeft bool) {
	if node == nil {
		return
	}
	for i := 0; i < depthOfTree; i++ {
		fmt.Print("|  ")
	}

	if depthOfTree > 0 {
		if isLeft {
			fmt.Print("|--")

		} else {
			fmt.Print("|__")
		}
	}

	fmt.Printf("%x\n", node.hash)

	displayMerkleTree(node.left, depthOfTree+1, true)
	displayMerkleTree(node.right, depthOfTree+1, false)
}

func (bc *BlockChain) DisplayBlockchain() {
	fmt.Println("Blockchain:")
	if bc.tail != nil {
		for block := bc.tail; block != nil; block = block.previous {
			fmt.Println("------------------------------------------------------------------------------------------------------------------")
			fmt.Printf("Transactions: %s \n Nonce: %d \n Hash: %x \n Previous Hash: %x\n Date and time of Transaction: %s \n", block.transactions, block.nonce, block.hash, block.previousHash, block.timeOfTransaction)
			fmt.Println("------------------------------------------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Println("Merkle Tree:")
			displayMerkleTree(block.merkleRoot, 0, false)
			fmt.Println("------------------------------------------------------------------------------------------------------------------")
		}
	} else {
		fmt.Println("Blockchain is empty")
	}
}

func (bc *BlockChain) ChangeBlock(transactionHash string, newTransaction string, blockHash string) {
	var transactionBlock *Block = nil
	for block := bc.tail; block != nil; block = block.previous {
		hash := fmt.Sprintf("%x", []byte(block.hash))
		if hash == blockHash {
			transactionBlock = block
			break
		}
	}

	merkleRoot := transactionBlock.merkleRoot
	transactionNode := findLeafNode(transactionHash, merkleRoot)
	transactionNode.hash = CalculateHash(newTransaction)
	fmt.Println("Transaction changed successfully")

}

func findLeafNode(transactionHash string, node *MerkleNode) *MerkleNode {
	//Check if we have reached a leaf node and if so return the node if its hash matches with the transaction hash
	if node.left == nil && node.right == nil {
		nodeHash := fmt.Sprintf("%x", []byte(node.hash))
		if nodeHash == transactionHash {
			return node
		} else {
			return nil
		}
	}
	//In case we have not reached a leaf node, move to the left node and check if its a leaf node
	leftNode := findLeafNode(transactionHash, node.left)
	if leftNode != nil {
		return leftNode
	}
	//In case we have not reached a leaf node, move to the right node and check if its a leaf node
	rightNode := findLeafNode(transactionHash, node.right)
	if rightNode != nil {
		return rightNode
	}
	//In case the transaction hash does not match with any of the leaf nodes' transactions, return nil indicating the transaction
	//is not present in the block
	return nil
}

func findNode(node *MerkleNode) *MerkleNode {
	//Check if we have reached a leaf node and if so return the node if its hash matches with the transaction hash
	if node.left == nil && node.right == nil {
		return node
	}
	//In case we have not reached a leaf node, move to the left node and check if its a leaf node
	leftNode := findNode(node.left)
	if leftNode != nil {
		return leftNode
	}
	//In case the transaction hash does not match with any of the leaf nodes' transactions, return nil indicating the transaction
	//is not present in the block
	return nil
}
func rebuildMerkleTree(nodes []*MerkleNode) string {
	for len(nodes) > 1 {

		var newNodes []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			hash := CalculateHash(fmt.Sprintf("%x", []byte(left.hash)) + fmt.Sprintf("%x", []byte(right.hash)))
			newNode := &MerkleNode{left, right, hash}
			newNodes = append(newNodes, newNode)

		}

		nodes = newNodes
	}

	return nodes[0].hash
}
func (bc *BlockChain) VerifyBlockchain() bool {
	if bc.tail != nil {
		for block := bc.tail; block != nil; block = block.previous {

			var node *MerkleNode = block.merkleRoot
			var merkleRootHash string = node.hash

			var nodes []*MerkleNode = GetLeftLeafNodes(node)

			var rootHash string = rebuildMerkleTree(nodes)

			if merkleRootHash != rootHash {
				fmt.Printf("The block with hash %x is modified.\n", block.hash)
				return false
			} else {
				continue
			}
		}
	} else {
		fmt.Println("Blockchain is empty")
		return true
	}
	return true
}

func GetLeftLeafNodes(node *MerkleNode) []*MerkleNode {

	if node == nil {
		return nil
	}

	if node.left == nil && node.right == nil {
		return []*MerkleNode{node}
	}

	leftNodes := GetLeftLeafNodes(node.left)
	rightNodes := GetLeftLeafNodes(node.right)

	return append(leftNodes, rightNodes...)
}

func findTraceToTransaction(transactionHash string, node *MerkleNode) []string {

	if node == nil {
		return nil
	}

	if node.left == nil && node.right == nil {
		hash := fmt.Sprintf("%x", []byte(node.hash))
		if hash == transactionHash {
			return []string{node.hash}
		} else {
			return nil
		}
	}

	leftTrace := findTraceToTransaction(transactionHash, node.left)
	if leftTrace != nil {
		return append(leftTrace, node.right.hash)
	}

	rightTrace := findTraceToTransaction(transactionHash, node.right)
	if rightTrace != nil {
		return append(rightTrace, node.left.hash)
	}

	return nil
}

func (bc *BlockChain) SetNumberOfTransactionsPerBlock(numberOfTransactionsPerBlock int) {
	bc.numberOfTransactionsPerBlock = numberOfTransactionsPerBlock
}
func (bc *BlockChain) SetBlockHashRange(min string, max string) {
	bc.minBlockHash = min
	bc.maxBlockHash = max
}

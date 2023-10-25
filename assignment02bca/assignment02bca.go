package assignment02bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type MerkleNode struct {
	left   *MerkleNode
	right  *MerkleNode
	parent *MerkleNode
	hash   string
}

type Block struct {
	transactions      []string
	nonce             int
	previousHash      string
	hash              string
	timeOfTransaction time.Time
	previous          *Block
	merkleRoot        *MerkleNode
}

type BlockChain struct {
	tail *Block
}

func CalculateHash(stringToHash string) string {
	h := sha256.New()
	h.Write([]byte(stringToHash))
	return string(h.Sum(nil))
}

func (b *Block) buildMerkleTree() *MerkleNode {
	var nodes []*MerkleNode

	for _, transaction := range b.transactions {
		nodes = append(nodes, &MerkleNode{nil, nil, nil, CalculateHash(transaction)})
	}

	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		var newNodes []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			hash := CalculateHash(left.hash + right.hash)

			newNode := &MerkleNode{left, right, nil, hash}

			left.parent = newNode
			right.parent = newNode
			newNodes = append(newNodes, newNode)

		}

		nodes = newNodes
	}

	return nodes[0]

}

func (bc *BlockChain) AddBlock(transactions []string, nonce int) {

	newBlock := &Block{transactions, nonce, "", "", time.Now(), bc.tail, nil}
	if bc.tail != nil {
		newBlock.previousHash = bc.tail.hash
	} else {

		newBlock.previousHash = ""
	}
	newBlock.merkleRoot = newBlock.buildMerkleTree()

	str := newBlock.merkleRoot.hash + strconv.Itoa(newBlock.nonce) + newBlock.previousHash + newBlock.timeOfTransaction.String()
	newBlock.hash = CalculateHash(str)

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
	for block := bc.tail; block != nil; block = block.previous {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("Transactions: %s \n Nonce: %d \n Hash: %x \n Previous Hash: %x\n Date and time of Transaction: %s \n", block.transactions, block.nonce, block.hash, block.previousHash, block.timeOfTransaction)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Println()
		fmt.Println("Merkle Tree:")
		displayMerkleTree(block.merkleRoot, 0, false)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
}

func findLeafNode(transactionHash string, node *MerkleNode) *MerkleNode {
	//Check if we have reached a leaf node and if so return the node if its hash matches with the transaction hash
	if node.left == nil && node.right == nil {
		if node.hash == transactionHash {
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

func (bc *BlockChain) VerifyTransaction(transactionHash string, blockHash string) bool {
	var transactionBlock *Block = nil
	for block := bc.tail; block != nil; block = block.previous {
		if block.hash == blockHash {
			transactionBlock = block
			break
		}
	}
	if transactionBlock != nil {
		merkleRoot := transactionBlock.merkleRoot
		merkleRootHash := merkleRoot.hash

		var pathToTransaction []string
		node := findLeafNode(transactionHash, merkleRoot)

		for node != nil {
			if node.left != nil && node.right != nil {
				if node.left.hash == transactionHash {
					pathToTransaction = append(pathToTransaction, node.right.hash)
				} else if node.right.hash == transactionHash {
					pathToTransaction = append(pathToTransaction, node.left.hash)
				}
			}
			node = node.parent
		}
		transacHash := transactionHash
		for _, hash := range pathToTransaction {
			transacHash = transacHash + hash
		}

		transacHash = transacHash + merkleRootHash

		computedMerkleRoot := CalculateHash(transacHash)

		if computedMerkleRoot == merkleRootHash {
			return true
		}

	}
	return false
}

package main

import (
	"fmt"
	"strings"
	"time"
)

/////////////////////
// Block
/////////////////////
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("timestamp         %d\n", b.timestamp)
	fmt.Printf("nonce             %d\n", b.nonce)
	fmt.Printf("previousHash      %s\n", b.previousHash)
	fmt.Printf("transactions      %s\n", b.transactions)
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp: time.Now().UnixNano(),
		nonce: nonce,
		previousHash: previousHash,
	}
}

/////////////////////
// BlockChain
/////////////////////
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
		fmt.Printf("%s \n", strings.Repeat("*", 25))
	}
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}


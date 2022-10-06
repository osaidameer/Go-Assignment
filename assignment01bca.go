package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct{
	hash_prev string
	hash_curr string
	transaction string
	nonce int 
	block_id int
}

type Blockchain struct{
	list []*Block
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

//func AddBlock

func (block_chain *Blockchain) NewBlock(transaction string, nonce int, previousHash string, id int) {
	newBlock := new(Block)
	newBlock.hash_prev = previousHash
	newBlock.transaction = transaction
	newBlock.block_id = id
	newBlock.nonce = nonce
	hash := previousHash + transaction + strconv.Itoa(nonce) + strconv.Itoa(id)
	newBlock.hash_curr = CalculateHash(hash)
	block_chain.list = append(block_chain.list, newBlock)
}

func (block_chain *Blockchain) ListBlock(){
	for i:=0; i<len(block_chain.list); i++ {
		fmt.Printf("\n%s Block %d %s\n", strings.Repeat("=", 36), i, strings.Repeat("=", 36))
		fmt.Println("Previous Hash: ", block_chain.list[i].hash_prev)
		fmt.Println("Current Hash: ", block_chain.list[i].hash_curr)
		fmt.Println("Nonce: ", block_chain.list[i].nonce)
		fmt.Println("Transaction: ", block_chain.list[i].transaction, "\n")
		fmt.Println("\t\t|\n\t\t|\n\t\t|\n\t\t----\n\t\t   \\/")
		//fmt.Println(block_chain.list[i].hash_curr)
	}
}

func (block_chain *Blockchain) ChangeBlock(id int, new_transaction string){
	block_chain.list[id].transaction = new_transaction
	hash := block_chain.list[id].hash_prev + block_chain.list[id].transaction + strconv.Itoa(block_chain.list[id].nonce) + strconv.Itoa(block_chain.list[id].block_id)
	block_chain.list[id].hash_curr = CalculateHash(hash)
}

func (block_chain *Blockchain) VerifyChain() (bool, int){
	for i:=0; i<len(block_chain.list); i++ {
		if(i!=0){
			hash := block_chain.list[i-1].hash_prev + block_chain.list[i-1].transaction + strconv.Itoa(block_chain.list[i-1].nonce) + strconv.Itoa(block_chain.list[i-1].block_id)
			check_hash := CalculateHash(hash)
			if check_hash != block_chain.list[i].hash_prev {
				return false, i-1
			}
		}
	}
	return true, 0
}

func getInput(prompt string, r *bufio.Scanner) (string) {
	fmt.Printf(prompt)
	r.Scan()
	return r.Text()
}


func menu(block_chain * Blockchain){
	block_num := 0
	choice:= 0
	for {
		fmt.Println("\n1) Enter 1 to create new block")
		fmt.Println("2) Enter 2 to list blockchain")
		fmt.Println("3) Enter 3 to change block")
		fmt.Println("4) Enter 4 to verify chain")
		fmt.Println("0) Enter 0 to quit")
		fmt.Printf("Enter choice: ")
		fmt.Scanln(&choice)
		switch choice{
		case 1:
			//var transaction string
			reader := bufio.NewScanner(os.Stdin)
			transaction := getInput("\nEnter Transaction: ", reader)
			nonce := rand.Intn(1000)
			prev_hash := block_chain.list[block_num].hash_curr
			block_num += 1
			block_chain.NewBlock(transaction, nonce, prev_hash, block_num)
			fmt.Println("\nNew Block Created!")
			break
		case 2:
			block_chain.ListBlock()
			break
		case 3:
			if block_num == 0 {
				fmt.Println("\nNo blocks to change!")
				break
			}
		    reader := bufio.NewScanner(os.Stdin)
			var block_id int
			fmt.Printf("\nNumber of Blocks: %d", len(block_chain.list))
			fmt.Printf("\nEnter block ID to change: ")
			fmt.Scanln(&block_id)
			if block_id > block_num || block_id < 0 {
				fmt.Println("\nBlock does not exist!")
				break
			} else if block_id == 0 {
				fmt.Println("\nCan not change genesis block!")
				break
			}
			transaction := getInput("\nEnter new transaction: ", reader)
			block_chain.ChangeBlock(block_id, transaction)
			fmt.Println("\nBlock Updated!")
			break
		case 4:
			test, index := block_chain.VerifyChain()
			if test == true{
				fmt.Println("\nBlockchain Verified!")
			} else {
				fmt.Printf("\nBlock %d has been tampered!", index)
				fmt.Println("\nBlockchain Tampered!")
			}
			break
		case 0:
			fmt.Println("\nExiting...")
			return
		default:
			fmt.Println("\nInvalid Input!")
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	block_chain := new(Blockchain)
	block_chain.NewBlock("Genesis Block", rand.Intn(1000), "", 0)
	menu(block_chain)
}
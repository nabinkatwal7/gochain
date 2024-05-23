package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

func (commandLineInterface *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (commandLineInterface *CLI) validateArgs() {
	if len(os.Args) < 2 {
		commandLineInterface.printUsage()
		os.Exit(1)
	}
}

func (commandLineInterface *CLI) addBlock(data string) {
	commandLineInterface.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (commandLineInterface *CLI) printChain() {
	blockChainIterator := commandLineInterface.bc.Iterator()

	for {
		block := blockChainIterator.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		proofOfWork := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (commandLineInterface *CLI) Run() {
	commandLineInterface.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		commandLineInterface.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		commandLineInterface.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		commandLineInterface.printChain()
	}
}

package main

func main() {
	blockChain := NewBlockChain()
	defer blockChain.db.Close()

	commandLineInterface := CLI{blockChain}
	commandLineInterface.Run()
}

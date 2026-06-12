package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gbx-operator ledger <ledger.jsonl>")
		return
	}

	switch os.Args[1] {
	case "ledger":
		path := "ledger.jsonl"
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		if err := printLedger(path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func printLedger(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

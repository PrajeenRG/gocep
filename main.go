package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	printHelp()
}

func printHelp() {
	fmt.Printf("gocep -- a cli tool to encrypt/decrypt data\n")
	fmt.Printf("\nUsage:\n\n")
	fmt.Printf("\t\tgocep command [options]\n")
	fmt.Printf("\nCommands:\n\n")
	fmt.Printf("\tencrypt\t\tPuts the tool into encryption mode\n")
	fmt.Printf("\tdecrypt\t\tPuts the tool into decryption mode\n")
	fmt.Printf("\nOptions:\n\n")
	fmt.Printf("\t--cipher, -c\tAllows user to choose encryption, not needed for decryption\n")
	fmt.Printf("\t--string, -s\tToggles option to give inline data\n")
	fmt.Printf("\t--file, -f  \tToggles file mode for data\n")
	fmt.Printf("\t--output, -o\tStore result into a file, prints it if not specified\n")
	fmt.Printf("\t--help, -h  \tPrints help message\n")
	fmt.Printf("\nExample:\n\n")
	fmt.Printf("\tgocep encrypt -c 1 -s use_this_text --output sample.txt\n")
	fmt.Printf("\tgocep decrypt -f sample.txt\n\n")
}

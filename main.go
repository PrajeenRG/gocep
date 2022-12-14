package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	printHelp()
}

func printHelp() {
	fmt.Println(strings.TrimSpace(
		"gocep - a command-line encryption/decryption tool\n" +
			"\n" +
			"Usage:\n" +
			"    gocep command [options] <arguments>\n" +
			"\n" +
			"Commands:\n" +
			"    encrypt			to select encryption mode\n" +
			"    decrypt			to select decryption mode\n" +
			"\n" +
			"Options:\n" +
			"    -c, --cipher	to set cipher algorithm to use\n" +
			"    -s, --string	pass data as a string\n" +
			"    -f, --file		pass data as a file\n" +
			"    -o, --output	destination file to store the result,\n" +
			"                    defaults to stdout for result display\n" +
			"    -h, --help		display help\n" +
			"\n" +
			"Examples:\n" +
			"    gocep encrypt -c 1 -s use_this_text --output sample.txt\n" +
			"    gocep decrypt -f sample.txt"))
}

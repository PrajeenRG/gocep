package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/prajeenrg/gocep/internal/cipher"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}
	if len(args)%2 == 0 {
		fmt.Println("Error!!! Invalid options have been specified.")
		fmt.Println("Check 'gocep -h' for more information.")
		return
	}

	argMap := decodeArgs(args)

	var result string

	switch argMap["mode"] {
	case "encrypt":
		result = processEncryption(argMap)
	case "decrypt":
		result = processDecryption(argMap)
	default:
		fmt.Println("Error!!! Invalid options have been specified.")
		fmt.Println("Check 'gocep -h' for more information.")
		return
	}

	if file, ok := argMap["output"]; ok {
		fmt.Println(result)
	} else {
		err := ioutil.WriteFile(file, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error!!! File not writable")
		} else {
			fmt.Println("Done!!! Output has been written to " + file)
		}
	}

}

func cipherProvider(code string) cipher.Cipher {
	var alg cipher.Cipher
	switch code {
	case "0":
		alg = cipher.Bitwise{}
	case "1":
		alg = cipher.Block{"128"}
	case "2":
		alg = cipher.Caesar{}
	case "3":
		alg = cipher.Stream{}
	case "4":
		alg = cipher.Transpose{5}
	case "5":
		alg = cipher.Vigenere{}
	case "6":
		alg = cipher.Xor{65}
	default:
		fmt.Println("Error!!!! Invalid cipher algorithm")
		os.Exit(1)
	}
	return alg
}

func processEncryption(argMap map[string]string) string {
	alg := cipherProvider(argMap["cipher"])
	return alg.Encrypt(argMap["data"])
}

func processDecryption(argMap map[string]string) string {
	var code string
	if val, ok := argMap["cipher"]; ok {
		code = val
	} else {
		code = string(argMap["data"][len(argMap["data"])-1])
	}
	alg := cipherProvider(code)
	return alg.Decrypt(argMap["data"])
}

func decodeArgs(args []string) map[string]string {
	argMap := make(map[string]string)
	argMap["mode"] = args[0]
	for i := 1; i < len(args); i += 2 {
		argMap[args[i]] = args[i+1]
	}
	return argMap
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

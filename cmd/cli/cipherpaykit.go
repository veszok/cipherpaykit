package main

import (
	"fmt"
	"os"
	"github.com/veszok/cipherpaykit/internal/kcv"
	"github.com/veszok/cipherpaykit/internal/tdes"
)

func showUsage() {
	fmt.Printf(`NAME:
   cipherpaykit - A tool providing cryptographic operations and payment validation

USAGE:
   cipherpaykit [command [subcommand]] [flags...]

COMMANDS:
   help    Shows a list of commands or help for one command
   tdes    Encrypts/Decrypts data using triple DES algorithm
   kcv     Calculates Key Check Value for the given key

SUBCOMMANDS:
   tdes:
      enc  Encrytps data using triple DES
	  dec  Decrypts data using triple DES
FLAGS:
   --key   Key data
   --mode  Block mode: ecb, cbc. Default: ecb
   --data  Data to be encrypted/decrypted
   --iv    Initialization vector (mandatory for CBC block mode)
`)

}

type cmdProcessMapType = map[string]func([]string) (string, error)

var cmdMap = cmdProcessMapType{
	"tdes": tdes.ProcessArgs,
	"kcv":  kcv.ProcessArgs,
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "help" {
		showUsage()
		return
	}
	processFunc, ok := cmdMap[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "Invalid subcommand %q\n", os.Args[1])
		os.Exit(1)
	}

	res, err := processFunc(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

package kcv

import (
	"encoding/hex"
	"flag"
	"os"

	"github.com/veszok/cipherpaykit"
)

type argType struct {
	key string
	// format string
}

// go run main.go kcv --key 010101010101010101010101010101010101010101010101
func (c *argType) process() (string, error) {
	key, err := hex.DecodeString(c.key)
	if err != nil {
		return "", err
	}
	res, err := cipherpaykit.GenerateKcv(key)
	if err != nil {
		return "", err
	}
	return res, nil
}

func ProcessArgs(args []string) (string, error) {

	var argData argType

	cmd := flag.NewFlagSet("kcv", flag.ExitOnError)
	cmd.StringVar(&argData.key, "key", "", "Key for KCV calculation")
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return "", err
	}

	return argData.process()
}

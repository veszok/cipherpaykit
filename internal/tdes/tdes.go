package tdes

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/veszok/cipherpaykit"
)

type opType int

const (
	opEncrypt opType = iota
	opDecrypt
)

type modeType int

const (
	modeEcb modeType = iota
	modeCbc
)

type argType struct {
	mode modeType
	key  string
	data string
	iv   string
	// format string
	op opType
}

func (c *argType) process() (string, error) {
	// assume format is hex
	key, err := hex.DecodeString(c.key)
	if err != nil {
		return "", err
	}
	data, err := hex.DecodeString(c.data)
	if err != nil {
		return "", err
	}
	var res []byte
	switch c.mode {
	case modeEcb:
		switch c.op {
		case opEncrypt:
			res, err = cipherpaykit.EncryptTripleDesEcb(key, data)
		case opDecrypt:
			res, err = cipherpaykit.DecryptTripleDesEcb(key, data)
		default:
			panic("invalid Op")
		}
	case modeCbc:
		if len(c.iv) == 0 {
			return "", errors.New("iv doesn't exist in CBC mode")
		}
		var iv []byte
		iv, err = hex.DecodeString(c.iv)
		if err != nil {
			return "", err
		}
		switch c.op {
		case opEncrypt:
			res, err = cipherpaykit.EncryptTripleDesCbc(key, data, iv)
		case opDecrypt:
			res, err = cipherpaykit.DecryptTripleDesCbc(key, data, iv)
		default:
			panic("invalid Op")
		}
	default:
		panic("invalid mode")
	}
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(res), nil
}

func ProcessArgs(args []string) (string, error) {

	// Current only support hex format

	var argData argType

	if len(os.Args) < 3 {
		return "", errors.New("TDES: no enc/dec subcommand specified")
	}
	switch os.Args[2] {
	case "enc":
		argData.op = opEncrypt
	case "dec":
		argData.op = opDecrypt
	default:
		return "", fmt.Errorf("TDES: Unknown subcommand %q", os.Args[2])
	}

	cmd := flag.NewFlagSet("tdes", flag.ExitOnError)
	mode := cmd.String("mode", "ecb", "Block mode for TDES")
	cmd.StringVar(&argData.key, "key", "", "Key for encryption/decryption")
	cmd.StringVar(&argData.data, "data", "", "Data to be encrypted/decrypted")
	cmd.StringVar(&argData.iv, "iv", "", "Initialization vector")
	err := cmd.Parse(os.Args[3:])
	if err != nil {
		return "", err
	}

	switch *mode {
	case "ecb":
		argData.mode = modeEcb
	case "cbc":
		argData.mode = modeCbc
	default:
		return "", fmt.Errorf("TDES: Unknown mode %q", *mode)
	}

	return argData.process()

}

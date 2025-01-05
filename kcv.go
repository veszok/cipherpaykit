// Package tripledes implements payment related utility functions.
package cipherpaykit

import (
	"encoding/hex"
)

// GenerateKcv generates the Key Check Value of the given key
func GenerateKcv(key []byte) (string, error) {
	res, err := EncryptTripleDesCbc(key, []byte("00000000"), []byte("00000000"))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(res)[:6], nil
}

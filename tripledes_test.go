package cipherpaykit

import (
	"encoding/hex"
	"testing"

	"github.com/veszok/cipherpaykit/internal/test"
)

func TestEncryptCbc(t *testing.T) {
	key := []byte("012345012345012345012345")
	data := []byte("aaaaaaaa")
	iv := []byte("00000000")
	res, _ := EncryptTripleDesCbc(key, data, iv)
	test.Assert(t, "b232c6dcf5873f26", hex.EncodeToString(res))
}

func TestDecryptCbc(t *testing.T) {
	key := []byte("012345012345012345012345")
	data, _ := hex.DecodeString("b232c6dcf5873f26")
	iv := []byte("00000000")
	res, _ := DecryptTripleDesCbc(key, data, iv)
	test.Assert(t, []byte("aaaaaaaa"), res)
}

func TestEncryptEcb(t *testing.T) {
	key := []byte("012345012345012345012345")
	data := []byte("aaaaaaaabbbbbbbb")
	res, _ := EncryptTripleDesEcb(key, data)
	test.Assert(t, "ba1e59ff04413cd902b85219a34d2cb4", hex.EncodeToString(res))
}

func TestDecryptEcb(t *testing.T) {
	key := []byte("012345012345012345012345")
	data, _ := hex.DecodeString("ba1e59ff04413cd902b85219a34d2cb4")
	res, _ := DecryptTripleDesEcb(key, data)
	test.Assert(t, []byte("aaaaaaaabbbbbbbb"), res)
}

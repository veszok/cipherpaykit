// Package tripledes implements convenient functions to encrypt/decrypt data using Triple DES algorithm.
package cipherpaykit

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func padKey(key []byte) ([]byte, error) {
	switch len(key) {
	case 24:
		return key, nil
	case 16:
		res := make([]byte, 24)
		copy(res, key)
		copy(res[16:], key[0:8])
		return res, nil
	case 8:
		res := make([]byte, 24)
		copy(res, key)
		copy(res[8:], key)
		copy(res[8:], key)
		return res, nil
	default:
		return nil, fmt.Errorf("key size (%d) should be 8, 16 or 24", len(key))
	}
}

// EncryptTripleDesCbc encrypts data with key and iv using Triple DES under CBC mode
func EncryptTripleDesCbc(inputKey []byte, data []byte, iv []byte) ([]byte, error) {
	key, err := padKey(inputKey)
	if err != nil {
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(fmt.Sprintf("Unexpected Error: %s", err.Error()))
	}
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("data length is not multiple of %d", blockSize)
	}
	res := make([]byte, len(data))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(res, data)
	return res, nil
}

// DecryptTripleDesCbc decrypts data with key and iv using Triple DES under CBC mode
func DecryptTripleDesCbc(inputKey []byte, data []byte, iv []byte) ([]byte, error) {
	key, err := padKey(inputKey)
	if err != nil {
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("data length is not multiple of %d", blockSize)
	}
	res := make([]byte, len(data))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(res, data)
	return res, nil
}

// EncryptTripleDesEcb encrypts data with key using Triple DES under ECB mode
func EncryptTripleDesEcb(inputKey []byte, data []byte) ([]byte, error) {
	key, err := padKey(inputKey)
	if err != nil {
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(fmt.Sprintf("Unexpected Error: %s", err.Error()))
	}
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("data length is not multiple of %d", blockSize)
	}
	res := make([]byte, len(data))

	for start := 0; start < len(res); start += blockSize {
		block.Encrypt(res[start:start+blockSize], data[start:start+blockSize])
	}
	return res, nil
}

// DecryptTripleEcb decrypts data with key using Triple DES under ECB mode
func DecryptTripleDesEcb(inputKey []byte, data []byte) ([]byte, error) {
	key, err := padKey(inputKey)
	if err != nil {
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		return nil, fmt.Errorf("data length is now multiple of %d", blockSize)
	}
	res := make([]byte, len(data))
	for start := 0; start < len(res); start += blockSize {
		block.Decrypt(res[start:start+blockSize], data[start:start+blockSize])
	}
	return res, nil
}

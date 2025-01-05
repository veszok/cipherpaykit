package cipherpaykit

import (
	"encoding/hex"
	"testing"

	"github.com/veszok/cipherpaykit/internal/test"
)

func TestGenerateKcv(t *testing.T) {
	key, _ := hex.DecodeString("010101010101010101010101010101010101010101010101")
	res, _ := GenerateKcv(key)
	test.Assert(t, "8ca64d", string(res))
}

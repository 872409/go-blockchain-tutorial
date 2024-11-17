package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// SHA256 hasing
func calculateHash2(block BlockV2) string {
	record := fmt.Sprintf("%d-%s-%s-%d-%d-%s", block.Index, block.PreviousHash, block.Data, block.Nonce, block.Timestamp, block.MinerId)
	//fmt.Println(record, block.Input, strings.Compare(record, block.Input))
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

type BlockV2 struct {
	Index        int    `json:"index"`
	PreviousHash string `json:"previousHash"`
	Data         string `json:"data"`
	Nonce        int    `json:"nonce"`
	Timestamp    int64  `json:"timestamp"`
	MinerId      string `json:"minerId"`
	Input        string `json:"input"`
	Hash         string `json:"hash"`
}

func Test_calchash(t *testing.T) {
	data := `{"index":1,"previousHash":"string","data":"string","nonce":999,"timestamp":1731725426047,"minerId":"string","input":"1-string-string-999-1731725426047-string","hash":"d0481602c3765fe69420743c6e601e92c35a64df5982b2fdd158818a3f528aae"}`

	var block BlockV2
	json.Unmarshal([]byte(data), &block)
	fmt.Println(block)

	hash := calculateHash2(block)
	fmt.Println(hash, block.Hash, strings.Compare(hash, block.Hash))

}

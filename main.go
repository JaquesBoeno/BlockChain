package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHah   string
}

var BlockChain []Block

func calculateHash(block Block) string {
	record := strings.Builder{}
	record.WriteString(fmt.Sprint(block.Index))
	record.WriteString(block.Timestamp)
	record.WriteString(fmt.Sprint(block.BPM))
	record.WriteString(block.PrevHah)

	h := sha256.New()
	h.Write([]byte(record.String()))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func main() {

}

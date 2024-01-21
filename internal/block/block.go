package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func calculateHash(block Block) string {
	record := strings.Builder{}
	record.WriteString(fmt.Sprint(block.Index))
	record.WriteString(block.Timestamp)
	record.WriteString(fmt.Sprint(block.BPM))
	record.WriteString(block.PrevHash)

	h := sha256.New()
	h.Write([]byte(record.String()))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = calculateHash(newBlock)

	return newBlock, nil
}

func IsValidBlock(newBlock, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

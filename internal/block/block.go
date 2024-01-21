package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type BlockInterface struct {
}

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func (b *BlockInterface) calculateHash(block Block) string {
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

func (b *BlockInterface) GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = b.calculateHash(newBlock)

	return newBlock, nil
}

func (b *BlockInterface) IsValidBlock(newBlock, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if b.calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

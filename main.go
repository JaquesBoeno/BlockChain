package main

import (
	"log"
	"time"

	Block "github.com/JaquesBoeno/BlockChain/internal/block"
	"github.com/JaquesBoeno/BlockChain/internal/router"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Block.BlockChain = append(Block.BlockChain, genesisBlock)
	}()

	log.Fatal(router.Run())
}

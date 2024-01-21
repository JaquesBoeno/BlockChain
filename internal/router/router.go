package router

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	Block "github.com/JaquesBoeno/BlockChain/internal/block"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

type Message struct {
	BPM int
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func run() error {
	mux := makeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("Listening on ", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Block.BlockChain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}

	defer r.Body.Close()

	newBlock, err := Block.GenerateBlock(Block.BlockChain[len(Block.BlockChain)-1], m.BPM)

	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if Block.IsValidBlock(newBlock, Block.BlockChain[len(Block.BlockChain)-1]) {
		newBlockChain := append(Block.BlockChain, newBlock)
		Block.ReplaceChain(newBlockChain)
		spew.Dump(Block.BlockChain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Mainchain Blockchain
var BClient BlockchainClient
var PORT int

func generateRandomPeerPort() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return 8000 + r.Intn(100)
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/blocks", returnAllBlocks).Methods("GET")
	r.HandleFunc("/api/mine", createBlock).Methods("POST")

	port := ":" + strconv.FormatInt(int64(PORT), 10)
	fmt.Printf("Listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func returnAllBlocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /api/blocks")
	json.NewEncoder(w).Encode(Mainchain)
}

func createBlock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST: /api/mine")
	body, _ := ioutil.ReadAll(r.Body)
	t := struct {
		Data string
	}{}

	err := json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
		return
	}
	if t.Data == "" {
		fmt.Println("No data supplied for block")
		return
	}

	Mainchain.addBlock(Data(t.Data))
	json.NewEncoder(w).Encode(Mainchain)
}

func main() {
	if len(os.Getenv("GENERATE_PEER_PORT")) > 0 {
		PORT = generateRandomPeerPort()
	} else {
		PORT = 8000
	}

	Mainchain = makeBlockchain()
	Mainchain.addBlock("This")
	Mainchain.addBlock("Is")
	Mainchain.addBlock("A")
	Mainchain.addBlock("Test")

	BClient = Client()
	go BClient.listen()

	BClient.publish(Mainchain)
	handleRequests()

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/blocks", returnAllBlocks)
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllBlocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBlocks")
	json.NewEncoder(w).Encode(Mainchain)
}

func main() {
	Mainchain = makeBlockchain()
	Mainchain.addBlock("This")
	Mainchain.addBlock("Is")
	Mainchain.addBlock("A")
	Mainchain.addBlock("Test")

	handleRequests()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homePage)
	r.HandleFunc("/api/blocks", returnAllBlocks).Methods("GET")
	r.HandleFunc("/api/mine", createBlock).Methods("POST")

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func returnAllBlocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBlocks")
	json.NewEncoder(w).Encode(Mainchain)
}

func createBlock(w http.ResponseWriter, r *http.Request) {
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
	Mainchain = makeBlockchain()
	Mainchain.addBlock("This")
	Mainchain.addBlock("Is")
	Mainchain.addBlock("A")
	Mainchain.addBlock("Test")

	handleRequests()
}

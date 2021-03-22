package main

import (
	"fmt"
)

func main() {
	b := makeGenesis()
	b2 := mineBlock(b, "data")

	fmt.Println(b)
	fmt.Println(b2)

}

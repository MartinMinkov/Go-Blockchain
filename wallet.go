package main

var STARTING_BALANCE = float64(1000)

type Wallet struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string
	Balance    float64 `json:"balance"`
}

func CreateWallet() Wallet {
	privateKey, publicKey := generateKeypair()
	return Wallet{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Balance:    STARTING_BALANCE,
	}
}

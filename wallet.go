package main

var STARTING_BALANCE = float64(1000)

type Wallet struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string
	Balance    float64 `json:"balance"`
}

func CreateWallet() Wallet {
	return Wallet{
		PublicKey:  "pub-key",
		PrivateKey: "priv-key",
		Balance:    STARTING_BALANCE,
	}
}

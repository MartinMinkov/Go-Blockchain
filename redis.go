package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var TEST_CHANNEL = "TEST_CHANNEL"
var BLOCKCHAIN_CHANNEL = "BLOCKCHAIN_CHANNEL"

type BlockchainClient struct {
	client  *redis.Client
	pubsub  *redis.PubSub
	channel string
}

func Client() BlockchainClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	channel := BLOCKCHAIN_CHANNEL
	pubsub := rdb.Subscribe(ctx, channel)
	return BlockchainClient{client: rdb, pubsub: pubsub, channel: channel}
}

func (b BlockchainClient) listen() {
	ch := b.pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		handleBlockchainPublish(msg.Payload)
	}
}

func (b BlockchainClient) publish(data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	b.client.Publish(ctx, b.channel, json)
}

func handleBlockchainPublish(payload string) {
	var newChain Blockchain
	bytes := []byte(payload)
	err := json.Unmarshal(bytes, &newChain)
	if err != nil {
		fmt.Println(err)
	}
	Mainchain = replaceBlockchain(Mainchain, newChain)
}

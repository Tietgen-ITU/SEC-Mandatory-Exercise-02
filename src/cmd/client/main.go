package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"math/rand"

	"google.golang.org/grpc"
	pb "sec.itu.dk/ex2/api"
	"sec.itu.dk/ex2/internals/commitments"
	"sec.itu.dk/ex2/internals/signatures"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:5001", "Server to play the dice game with")
	signatureHandler = signatures.CreateNew()
	commitmentHandler = commitments.CreateNew()
)

func main() {

	flag.Parse()

	fmt.Println("Setting up client...")
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("Could not connect to server!")
	}
	
	defer conn.Close()

	client := pb.NewDiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if pk, ok := exchangePublicKey(signatureHandler.PublicKey(), conn, &ctx); ok {

		playDiceGame(&pk, &client, &ctx)
	}
}

/*
Is the main game loop for the client
*/
func playDiceGame(pk *big.Int,client *pb.DiceClient, ctx *context.Context) {

	for i := 0; i < 6; i++ {
		// 1. Generate dice roll for client
		// 2. Generate dice roll for server
		// 3. Who won?
	}
}

func diceRoll(pk *big.Int, client pb.DiceClient, ctx *context.Context) (int, bool) {

	commitmentKey := rand.Int63()
	roll := rand.Int31n(6)
	for roll == 0 {
		roll = rand.Int31n(6)
	}

	commit := commitmentHandler.Commit(*big.NewInt(int64(roll)), *big.NewInt(commitmentKey))

	serverCommitment, err := client.Commit(*ctx, &pb.Commitment{
		Value: commit.Int64(),
	})

	if err != nil {
		fmt.Println("Could not send commitment to server!")
		return 0, false
	}

	serverRollReveal, err := client.Reveal(*ctx, &pb.CommitmentReveal{
		Value: roll,
		Key: commitmentKey,
		Signature: &pb.Signature{
			Signature: 0,
			Random: 0,
		},
	})

	if err != nil {
		fmt.Println("Could not reveal commitment to server!")
		return 0, false
	}

	serverValue := big.NewInt(int64(serverRollReveal.GetValue()))
	serverCommitmentKey := big.NewInt(serverRollReveal.GetKey())
	serverCommitmentValue := big.NewInt(serverCommitment.GetValue())

	correctMessage := commitmentHandler.Verify(*serverValue, *serverCommitmentValue, *serverCommitmentKey)

	result := serverRollReveal.Value ^ roll

	return int(result), correctMessage
}

/*
Exchanges the public key with the server. Simulates the PKI
*/
func exchangePublicKey(publicKey big.Int, conn *grpc.ClientConn, ctx *context.Context) (big.Int, bool) {

	exchangeClient := pb.NewKeyExchangeClient(conn)

	reply, err := exchangeClient.ExchangePk(*ctx, &pb.Key{
		Value: publicKey.Int64(),
	})

	if err != nil {
		fmt.Println("Could not exchange keys!")
		return *new(big.Int), false
	}

	return *big.NewInt(reply.Value), true
}

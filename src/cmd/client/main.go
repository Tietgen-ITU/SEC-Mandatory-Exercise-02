package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strconv"

	"google.golang.org/grpc"
	pb "sec.itu.dk/ex2/api"
	"sec.itu.dk/ex2/internals/commitments"
	"sec.itu.dk/ex2/internals/crypto/hashing"
	"sec.itu.dk/ex2/internals/utils"
)

var (
	serverAddr        = flag.String("serverAddr", "localhost:5001", "Server to play the dice game with")
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

	playDiceGame(&client, &ctx)
}

/*
Is the main game loop for the client
*/
func playDiceGame(client *pb.DiceClient, ctx *context.Context) {

	for i := 0; i < 6; i++ {

		if clientRoll, clientRollOk := diceRoll(*client, ctx); clientRollOk {

			if serverRoll, serverRollOk := diceRoll(*client, ctx); serverRollOk {

				utils.PrintDiceRollWinner(clientRoll, serverRoll)
			}
		}
	}
}

/*
Makes the random dice roll in collaboration with the server
*/
func diceRoll(client pb.DiceClient, ctx *context.Context) (utils.DiceRoll, bool) {

	// Generate commitment key and make dice roll
	commitmentKey := hashing.GenerateRandomByteArray()
	roll := rand.Int31n(6)
	for roll == 0 {
		roll = rand.Int31n(6)
	}

	// Send commitment of roll
	commit := commitmentHandler.Commit([]byte(strconv.Itoa(int(roll))), commitmentKey)
	serverCommitment, err := client.Commit(*ctx, &pb.Commitment{
		Value: commit,
	})

	if err != nil {
		fmt.Println("Could not send commitment to server!")
		return 0, false
	}

	// Reveal commitment to server
	serverRollReveal, err := client.Reveal(*ctx, &pb.CommitmentReveal{
		Value: roll,
		Key:   commitmentKey,
	})

	if err != nil {
		fmt.Println("Could not reveal commitment to server!")
		return 0, false
	}

	// Calculate the random roll
	serverValue := serverRollReveal.GetValue()
	serverCommitmentKey := serverRollReveal.GetKey()
	serverCommitmentValue := serverCommitment.Value

	correctMessage := commitmentHandler.Verify([]byte(strconv.Itoa(int(serverValue))), serverCommitmentValue, serverCommitmentKey)

	result := utils.CalculateRoll(utils.PartialRoll(serverRollReveal.Value), utils.PartialRoll(roll))

	return result, correctMessage
}
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	fmt.Printf("Server endpoint address from flag: %s \n", *serverAddr)

	fmt.Println("Setting up client...")

	// Get TLS credentials
	tlsCreds, err := getTLSCredentials()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create connection with TLS credentials
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(tlsCreds), grpc.WithChainUnaryInterceptor(), grpc.WithChainStreamInterceptor())
	if err != nil {
		fmt.Println("Could not connect to server!")
	}

	defer conn.Close()

	// Create Dice client
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
	roll := utils.RollPartially()

	// Send commitment of roll
	commit := commitmentHandler.Commit([]byte(strconv.Itoa(int(roll))), commitmentKey)
	serverCommitment, err := client.Commit(*ctx, &pb.Commitment{
		Value: commit,
	})

	if err != nil {
		fmt.Printf("Could not send commitment to server: %s \n", err.Error())
		return 0, false
	}

	// Reveal commitment to server
	serverRollReveal, err := client.Reveal(*ctx, &pb.CommitmentReveal{
		Value: int32(roll),
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

/*
Gets certificates and loads them into an application created certificate pool and returns the TLS credentials
*/
func getTLSCredentials() (credentials.TransportCredentials, error) {

	caCert, err := ioutil.ReadFile("./assets/certificates/ca-cert.crt")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(caCert) // Get PEM decoded

	// Parse certificates and add it to the certificate pool
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)

	// Create TLS configuration
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

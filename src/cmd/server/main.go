package main

import (
	"context"
	"errors"
	"math/big"
	"math/rand"

	"sec.itu.dk/ex2/internals/utils"
	pb "sec.itu.dk/ex2/api"
	"sec.itu.dk/ex2/internals/commitments"
	"sec.itu.dk/ex2/internals/signatures"
)

var (
	commitmentHandler = commitments.CreateNew()
	signatureHandler  = signatures.CreateNew()
)

const RESET_VALUE int = -1

type Server struct {
	pb.UnimplementedKeyExchangeServer
	pb.UnimplementedDiceServer
	clientRoll      utils.DiceRoll
	clientCommit    big.Int
	commitmentValue utils.PartialRoll
	commitmentKey   int64
	clientPk        big.Int
}

func main() {

}

func (s *Server) Commit(ctx context.Context, commit *pb.Commitment) (*pb.Commitment, error) {

	// TODO: Verify signature
	// TODO: Decrypt message

	s.clientCommit = *big.NewInt(commit.GetValue())

	s.commitmentKey = rand.Int63()
	s.commitmentValue = utils.PartialRoll(rand.Int31n(6))
	for s.commitmentValue == 0 {
		s.commitmentValue = utils.PartialRoll(rand.Int31n(6))
	}

	c := commitmentHandler.Commit(*big.NewInt(int64(s.commitmentValue)), *big.NewInt(s.commitmentKey))

	return &pb.Commitment{
		Value: c.Int64(),
	}, nil
}

func (s *Server) Reveal(ctx context.Context, reveal *pb.CommitmentReveal) (*pb.CommitmentReveal, error) {

	clientRoll := big.NewInt(int64(reveal.GetValue()))
	clientCommitmentKey := big.NewInt(reveal.GetKey())
	correctMessage := commitmentHandler.Verify(*clientRoll, s.clientCommit, *clientCommitmentKey)

	if !correctMessage {

		return nil, errors.New("not correct message")
	}

	if s.clientRoll == utils.DiceRoll(RESET_VALUE) {

		s.clientRoll = utils.CalculateRoll(utils.PartialRoll(reveal.GetValue()), s.commitmentValue)
	} else {

		defer s.resetRoll()
		serverRoll := utils.CalculateRoll(utils.PartialRoll(reveal.Value), s.commitmentValue)
		utils.PrintDiceRollWinner(s.clientRoll, serverRoll)
	}

	// Clean commitment values
	defer s.resetCommitment()

	return &pb.CommitmentReveal{
		Value: int32(s.commitmentValue),
		Key: s.commitmentKey,
		Signature: &pb.Signature{
			Signature: 0,
			Random: 0,
		},
	}, nil
}

func (s *Server) resetCommitment() {

	s.commitmentKey = int64(RESET_VALUE)
	s.commitmentValue = utils.PartialRoll(RESET_VALUE)
	s.clientCommit = *big.NewInt(int64(RESET_VALUE));
}

func (s *Server) resetRoll() {

	s.clientRoll = utils.DiceRoll(RESET_VALUE);
}

func (s *Server) ExchangePk(ctx context.Context, key *pb.Key) (*pb.Key, error) {

	s.clientPk = *big.NewInt(key.GetValue())
	pk := signatureHandler.PublicKey()

	return &pb.Key{
		Value: pk.Int64(),
	}, nil
}
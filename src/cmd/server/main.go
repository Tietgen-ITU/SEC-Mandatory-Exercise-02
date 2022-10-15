package main

import (
	"context"
	"errors"
	"math/rand"
	"strconv"

	pb "sec.itu.dk/ex2/api"
	"sec.itu.dk/ex2/internals/commitments"
	"sec.itu.dk/ex2/internals/crypto/hashing"
	"sec.itu.dk/ex2/internals/utils"
)

var (
	commitmentHandler = commitments.CreateNew()
)

const RESET_VALUE int = -1

type Server struct {
	pb.UnimplementedDiceServer
	clientRoll      utils.DiceRoll
	clientCommit    []byte
	commitmentValue utils.PartialRoll
	commitmentKey   []byte
}

func main() {

}

func (s *Server) Commit(ctx context.Context, commit *pb.Commitment) (*pb.Commitment, error) {

	s.clientCommit = commit.GetValue()

	s.commitmentKey = hashing.GenerateRandomByteArray()
	s.commitmentValue = utils.PartialRoll(rand.Int31n(6))
	for s.commitmentValue == 0 {
		s.commitmentValue = utils.PartialRoll(rand.Int31n(6))
	}

	c := commitmentHandler.Commit([]byte(strconv.Itoa(int(s.commitmentValue))), s.commitmentKey)

	return &pb.Commitment{
		Value: c,
	}, nil
}

func (s *Server) Reveal(ctx context.Context, reveal *pb.CommitmentReveal) (*pb.CommitmentReveal, error) {

	clientRoll := reveal.GetValue()
	clientCommitmentKey := reveal.GetKey()
	correctMessage := commitmentHandler.Verify([]byte(strconv.Itoa(int(clientRoll))), s.clientCommit, clientCommitmentKey)

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
	}, nil
}

func (s *Server) resetCommitment() {

	s.commitmentKey = nil
	s.commitmentValue = utils.PartialRoll(RESET_VALUE)
	s.clientCommit = nil
}

func (s *Server) resetRoll() {

	s.clientRoll = utils.DiceRoll(RESET_VALUE);
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiceClient is the client API for Dice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiceClient interface {
	Commit(ctx context.Context, in *Commitment, opts ...grpc.CallOption) (*Commitment, error)
	Reveal(ctx context.Context, in *CommitmentReveal, opts ...grpc.CallOption) (*CommitmentReveal, error)
}

type diceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiceClient(cc grpc.ClientConnInterface) DiceClient {
	return &diceClient{cc}
}

func (c *diceClient) Commit(ctx context.Context, in *Commitment, opts ...grpc.CallOption) (*Commitment, error) {
	out := new(Commitment)
	err := c.cc.Invoke(ctx, "/Dice/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diceClient) Reveal(ctx context.Context, in *CommitmentReveal, opts ...grpc.CallOption) (*CommitmentReveal, error) {
	out := new(CommitmentReveal)
	err := c.cc.Invoke(ctx, "/Dice/Reveal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiceServer is the server API for Dice service.
// All implementations must embed UnimplementedDiceServer
// for forward compatibility
type DiceServer interface {
	Commit(context.Context, *Commitment) (*Commitment, error)
	Reveal(context.Context, *CommitmentReveal) (*CommitmentReveal, error)
	mustEmbedUnimplementedDiceServer()
}

// UnimplementedDiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiceServer struct {
}

func (UnimplementedDiceServer) Commit(context.Context, *Commitment) (*Commitment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedDiceServer) Reveal(context.Context, *CommitmentReveal) (*CommitmentReveal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reveal not implemented")
}
func (UnimplementedDiceServer) mustEmbedUnimplementedDiceServer() {}

// UnsafeDiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiceServer will
// result in compilation errors.
type UnsafeDiceServer interface {
	mustEmbedUnimplementedDiceServer()
}

func RegisterDiceServer(s grpc.ServiceRegistrar, srv DiceServer) {
	s.RegisterService(&Dice_ServiceDesc, srv)
}

func _Dice_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commitment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dice/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiceServer).Commit(ctx, req.(*Commitment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dice_Reveal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitmentReveal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiceServer).Reveal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dice/Reveal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiceServer).Reveal(ctx, req.(*CommitmentReveal))
	}
	return interceptor(ctx, in, info, handler)
}

// Dice_ServiceDesc is the grpc.ServiceDesc for Dice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dice",
	HandlerType: (*DiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Commit",
			Handler:    _Dice_Commit_Handler,
		},
		{
			MethodName: "Reveal",
			Handler:    _Dice_Reveal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dice.proto",
}

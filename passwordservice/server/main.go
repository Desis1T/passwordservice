package main

import (
	"context"
	"log"
	"net"
	ps "passgen/passwordservice/proto"

	garbler "github.com/michaelbironneau/garbler/lib"
	"google.golang.org/grpc"
)

type PasswordGeneratorServiceServer struct {
}

// mustEmbedUnimplementedPasswordGeneratorServiceServer implements passwordservice.PasswordGeneratorServiceServer.
func (*PasswordGeneratorServiceServer) mustEmbedUnimplementedPasswordGeneratorServiceServer() {
	panic("unimplemented")
}

func (s *PasswordGeneratorServiceServer) Generate(ctx context.Context,
	req *ps.PasswordRequest) (*ps.PasswordResponse, error) {

	var err error
	response := new(ps.PasswordResponse)

	requirements := garbler.MakeRequirements(req.Sample)
	response.Password, err = garbler.NewPassword(&requirements)

	return response, err
}

func main() {
	server := grpc.NewServer()

	instance := new(PasswordGeneratorServiceServer)

	ps.RegisterPasswordGeneratorServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}

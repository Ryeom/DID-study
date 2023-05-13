package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ssikr/actors/issuer"
	"ssikr/config"
	protos "ssikr/protos"
)

func main() {

	// New Issuer
	issr := new(issuer.Issuer)
	issr.GenerateDID() //key 생성, DID, DID Document, VDR에 등록

	lis, err := net.Listen("tcp", config.SystemConfig.IssuerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	issuerServer := issuer.Server{}
	issuerServer.Issuer = issr

	s := grpc.NewServer()
	protos.RegisterSimpleIssuerServer(s, &issuerServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

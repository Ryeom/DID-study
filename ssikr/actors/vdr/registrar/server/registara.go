package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"ssikr/config"
	"ssikr/protos"

	"github.com/syndtr/goleveldb/leveldb"
)

// 정해진 형식! 구조체를 만들고 구현하려는 서버의 UnimplementedRegistrarServer라는 속성의 서버를 준다
type registrarServer struct {
	protos.UnimplementedRegistrarServer
}

func (server *registrarServer) RegisterDid(ctx context.Context, req *protos.RegistrarRequest) (*protos.RegistrarResponse, error) {
	log.Printf("Register DID: %s\n", req.Did)

	db, err := leveldb.OpenFile("did_db/dids", nil)
	if err != nil {
		//panic(/)
	}
	defer db.Close()

	err = db.Put([]byte(req.Did), []byte(req.DidDocument), nil)

	return &protos.RegistrarResponse{Result: "OK"}, nil
}

func main() {
	fmt.Println("### Start Registrar ###")
	lis, err := net.Listen("tcp", config.SystemConfig.RegistrarAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := registrarServer{}
	s := grpc.NewServer()
	protos.RegisterRegistrarServer(s, &server)

	log.Printf("Registrar Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"log"
	"net"

	"github.com/AirtonL/Klever-Crypto/database"
	"github.com/AirtonL/Klever-Crypto/pb"
	"github.com/AirtonL/Klever-Crypto/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	collection     *mongo.Collection
	serverPb       pb.CryptoProtoServer
	contextMongoDB context.Context
)

type CryptoProtoServer struct {
	pb.UnimplementedCryptoProtoServer
}

func main() {

	database.Init()

	contextMongoDB = database.GetContextMongo()
	collection = database.GetCollectionName("crypto")
	serverPb = services.NewCryptoServer(collection, contextMongoDB)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	servergRPC := grpc.NewServer()
	pb.RegisterCryptoProtoServer(servergRPC, serverPb)
	reflection.Register(servergRPC)

	log.Printf("Starting server on port: %v", lis.Addr())
	if err := servergRPC.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server stopped")

	database.MongoDisconect()

}

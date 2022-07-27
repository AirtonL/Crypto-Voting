package services

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/AirtonL/Klever-Crypto/database"
	"github.com/AirtonL/Klever-Crypto/pb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type Crypto struct {
	id         string
	name       string
	upvote     int64
	downvote   int64
	totalscore int64
}

const bufSize = 1024 * 1024

var (
	lis *bufconn.Listener

	service    pb.CryptoProtoServer
	mongoCtx   context.Context
	collection *mongo.Collection
)

func init() {
	database.Init()

	mongoCtx = database.GetContextMongo()
	collection = database.GetCollectionName("crypto_test")

	service = NewCryptoServer(collection, mongoCtx)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCryptoProtoServer(s, service)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateNewCrypto(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewCryptoProtoClient(conn)

	req := &pb.CreateNewCryptoRequest{
		Crypto: &pb.Crypto{
			Id:         "62de22ae8501922b4078bd8a",
			Name:       "TestCrypto",
			Upvote:     4,
			Downvote:   2,
			Totalscore: 2,
		},
	}

	result, err := client.CreateNewCrypto(ctx, req)

	if err != nil {
		t.Fatalf("Failed to create new crypto: %v", err)
	}

	log.Printf("Created new crypto: %v", result)

	assert.Equal(t, result.Crypto.Name, "TestCrypto")
}

// func TestGetByIdCrypto(t *testing.T) {

// 	ctx := context.Background()

// 	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatalf("Failed to dial bufnet: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewCryptoProtoClient(conn)

// 	mock := Crypto{
// 		id:         "62de22ae8501922b4078bd8b",
// 		name:       "TestCrypto",
// 		upvote:     4,
// 		downvote:   2,
// 		totalscore: 2,
// 	}

// 	req := &pb.GetByIdCryptoRequest{
// 		Id: "62de22ae8501922b4078bd8b",
// 	}

// 	result, err := client.GetByIdCrypto(ctx, req)

// 	if err != nil {
// 		t.Fatalf("Failed to get crypto by id: %v", err)
// 	}

// 	log.Printf("Successfully got crypto by id: %v", result.Crypto.Id)

// 	assert.Equal(t, result.Crypto.Name, mock.name)
// }

func TestUpVoteCrypto(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewCryptoProtoClient(conn)

	req := &pb.UpVoteCryptoRequest{
		Id: "62de22ae8501922b4078bd8a",
	}

	result, err := client.UpVoteCrypto(ctx, req)

	if err != nil {
		t.Fatalf("Failed to upvote crypto: %v", err)
	}

	log.Printf("Upvoted crypto: %v", result)

	assert.Equal(t, result.Upvote, true)
}

func TestDownVoteCrypto(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewCryptoProtoClient(conn)

	req := &pb.DownVoteCryptoRequest{
		Id: "62de22ae8501922b4078bd8a",
	}

	result, err := client.DownVoteCrypto(ctx, req)

	if err != nil {
		t.Fatalf("Failed to downvote crypto: %v", err)
	}

	log.Printf("Downvoted crypto: %v", result)

	assert.Equal(t, result.Downvote, true)
}

func TestDeleteCrypto(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewCryptoProtoClient(conn)

	req := &pb.DeleteCryptoRequest{
		Id: "62de22ae8501922b4078bd8a",
	}

	result, err := client.DeleteCrypto(ctx, req)

	if err != nil {
		t.Fatalf("Failed to delete crypto: %v", err)
	}

	log.Printf("Successfully deleted crypto")

	assert.Equal(t, result.Delete, true)
}

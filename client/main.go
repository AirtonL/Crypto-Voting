package main

import (

	// "io"
	"log"

	"github.com/AirtonL/Klever-Crypto/models"
	"github.com/AirtonL/Klever-Crypto/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	client pb.CryptoProtoClient
	crypto models.Cryptos
)

// https://www.youtube.com/watch?v=Y92WWaZJl24&ab_channel=TensorProgramming

func main() {
	log.Println("Starting client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	client := pb.NewCryptoProtoClient(connection)

	defer connection.Close()

	g := gin.Default()

	// request := pb.GetByIdCryptoRequest{
	// 	Id: "62df6577622d144887f96338",
	// }

	// responseId, err := client.GetByIdCrypto(context.Background(), &request)

	// if err != nil {
	// 	log.Fatalf("Could not get crypto: %v", err)
	// }

	// log.Printf("Crypto: %v", responseId)

	// ----------------------------------------------

	// request := pb.GetAllCryptoRequest{}

	// responseStream, err := client.GetAllCrypto(context.Background(), &request)
	// if err != nil {
	// 	log.Fatalf("Could not get cryptos: %v", err)
	// }

	// for {
	// 	stream, err := responseStream.Recv()
	// 	if err == io.EOF {
	// 		log.Fatalf("Could not receive response: %v", err)
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatalf("Could not get cryptos: %v", err)
	// 	}

	// 	log.Printf("crypto %v", stream.GetCrypto())
	// }
	// ------------------------------------------------
	// request := &pb.CreateNewCryptoRequest{
	// 	Crypto: &pb.Crypto{
	// 		Name:       "TestCrypto",
	// 		Upvote:     4,
	// 		Downvote:   2,
	// 		Totalscore: 2,
	// 	},
	// }
	// log.Printf("Connected to server %v", request)

	// res, err := client.CreateNewCrypto(context.Background(), request)

	// if err != nil {
	// 	log.Fatalf("Could not create new crypto: %v", err)
	// }

	// log.Printf("Created new crypto: %v", res.Crypto)
}

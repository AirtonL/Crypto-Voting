package main

import (

	// "io"

	"io"
	"log"
	"net/http"

	"github.com/AirtonL/Klever-Crypto/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	client pb.CryptoProtoClient
)

func main() {
	log.Println("Starting client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer connection.Close()
	log.Println(connection)

	client = pb.NewCryptoProtoClient(connection)

	g := gin.Default()

	g.POST("/crypto", createNewCrypto)
	g.GET("/crypto", getAllCrypto)
	g.PUT("/upvote/:id", upvoteCrypto)
	g.PUT("/downvote/:id", downvoteCrypto)
	g.GET("/crypto/:id", getByIdCrypto)
	g.DELETE("/crypto/:id", deleteCrypto)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

type nameBind struct {
	Name string `json:"name"`
}

func createNewCrypto(ctx *gin.Context) {
	var requestBody nameBind

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	req := pb.Crypto{
		Name: requestBody.Name,
	}

	log.Print(req)

	res, err := client.CreateNewCrypto(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func getAllCrypto(ctx *gin.Context) {
	obj := pb.GetAllCryptoRequest{}

	stream, err := client.GetAllCrypto(ctx, &obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func upvoteCrypto(ctx *gin.Context) {
	uid := ctx.Param("id")

	obj := pb.UpVoteCryptoRequest{Id: uid}

	res, err := client.UpVoteCrypto(ctx, &obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, res)
}

func downvoteCrypto(ctx *gin.Context) {
	uid := ctx.Param("id")

	obj := pb.DownVoteCryptoRequest{Id: uid}

	res, err := client.DownVoteCrypto(ctx, &obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, res)
}

func getByIdCrypto(ctx *gin.Context) {
	uid := ctx.Param("id")

	obj := pb.GetByIdCryptoRequest{Id: uid}

	res, err := client.GetByIdCrypto(ctx, &obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, res)
}

func deleteCrypto(ctx *gin.Context) {
	uid := ctx.Param("id")

	obj := pb.DeleteCryptoRequest{Id: uid}

	res, err := client.DeleteCrypto(ctx, &obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, res)
}

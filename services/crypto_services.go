package services

import (
	"context"
	"fmt"
	"log"

	"github.com/AirtonL/Klever-Crypto/models"
	"github.com/AirtonL/Klever-Crypto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cryptoServiceServer struct {
	collection     *mongo.Collection
	contextMongoDB context.Context
}

func NewCryptoServer(collection *mongo.Collection, contextMongoDB context.Context) pb.CryptoProtoServer {
	return &cryptoServiceServer{
		collection:     collection,
		contextMongoDB: contextMongoDB,
	}
}

func (m *cryptoServiceServer) CreateNewCrypto(ctx context.Context, in *pb.CreateNewCryptoRequest) (*pb.CreateCryptoResponse, error) {
	req := in.Crypto
	if req.GetName() == "" {
		return nil, fmt.Errorf("Crypto name is required")
	}

	requestBody := models.Crypto{
		Name:       req.GetName(),
		Upvote:     req.GetUpvote(),
		Downvote:   req.GetDownvote(),
		Totalscore: req.GetTotalscore(),
	}

	result, err := m.collection.InsertOne(m.contextMongoDB, requestBody)

	if err != nil {
		return nil, fmt.Errorf("Error creating new crypto: %v", err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	req.Id = oid.Hex()
	return &pb.CreateCryptoResponse{Crypto: req}, nil
}

func (m *cryptoServiceServer) GetByIdCrypto(ctx context.Context, in *pb.GetByIdCryptoRequest) (*pb.GetByIdCryptoResponse, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, fmt.Errorf("Error getting crypto by id: %v", err)
	}
	log.Println(oid)

	result := m.collection.FindOne(ctx, bson.M{"_id": oid})
	data := models.Crypto{}
	if err := result.Decode(&data); err != nil {
		return nil, fmt.Errorf("Error getting crypto by id: %v", err)
	}

	return &pb.GetByIdCryptoResponse{
		Crypto: &pb.Crypto{
			Id:         oid.Hex(),
			Name:       data.Name,
			Upvote:     data.Upvote,
			Downvote:   data.Downvote,
			Totalscore: data.Totalscore,
		}}, nil
}

func (m *cryptoServiceServer) GetAllCrypto(req *pb.GetAllCryptoRequest, stream pb.CryptoProto_GetAllCryptoServer) error {

	data := &models.Crypto{}

	cursor, err := m.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return fmt.Errorf("Error getting all crypto: %v", err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)

		if err != nil {
			return fmt.Errorf("Error getting all crypto: %v", err)
		}
		log.Println(data.Upvote)

		stream.Send(&pb.GetAllCryptoResponse{
			Crypto: &pb.Crypto{
				Id:         data.ID.Hex(),
				Name:       data.Name,
				Upvote:     data.Upvote,
				Downvote:   data.Downvote,
				Totalscore: data.Totalscore,
			},
		})
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("Error getting all crypto: %v", err)
	}
	return nil
}

func (m *cryptoServiceServer) UpVoteCrypto(ctx context.Context, in *pb.UpVoteCryptoRequest) (*pb.UpVoteCryptoResponse, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, fmt.Errorf("Error getting crypto by id: %v", err)
	}

	filter := bson.M{"_id": oid}

	_, err = m.collection.UpdateOne(ctx, filter, bson.M{"$inc": bson.M{"upvote": 1, "totalscore": 1}}, options.Update().SetUpsert(true))
	log.Println(err)
	if err != nil {
		return nil, fmt.Errorf("Error upvoting crypto: %v", err)
	}
	log.Println(filter)

	return &pb.UpVoteCryptoResponse{
		Upvote: true,
	}, nil
}

func (m *cryptoServiceServer) DownVoteCrypto(ctx context.Context, in *pb.DownVoteCryptoRequest) (*pb.DownVoteCryptoResponse, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, fmt.Errorf("Error getting crypto by id: %v", err)
	}

	filter := bson.M{"_id": oid}

	_, err = m.collection.UpdateOne(ctx, filter, bson.M{"$inc": bson.M{"downvote": 1, "totalscore": -1}}, options.Update().SetUpsert(true))

	if err != nil {
		return nil, fmt.Errorf("Error downvoting crypto: %v", err)
	}

	return &pb.DownVoteCryptoResponse{
		Downvote: true,
	}, nil
}

func (m *cryptoServiceServer) DeleteCrypto(ctx context.Context, in *pb.DeleteCryptoRequest) (*pb.DeleteCryptoResponse, error) {
	oid, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, fmt.Errorf("Error getting crypto by id: %v", err)
	}

	filter := bson.M{"_id": oid}

	_, err = m.collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("Error deleting crypto: %v", err)
	}

	return &pb.DeleteCryptoResponse{
		Delete: true,
	}, nil
}

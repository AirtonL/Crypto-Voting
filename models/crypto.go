package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Crypto struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name" binding:"required"`
	Upvote     int64              `bson:"upvote"`
	Downvote   int64              `bson:"downvote"`
	Totalscore int64              `bson:"totalscore"`
}

type Cryptos []Crypto

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	PostID                 primitive.ObjectID `bson:"_id"`
	UserId                 primitive.ObjectID `bson:"userId"`
	UserLatitude           string             `bson:"latitude"`
	UserLogitude           string             `bson:"longitude"`
	PostContent            string             `bson:"post_text"`
	PostBackgroundImageURL string             `bson:"post_bg_image,omitempty"`
	PostTages              []interface{}      `bson:"post_tags,omitempty"`
	PostedTime             time.Time          `bson:"posted_time"`
	PostLikedList          []interface{}      `bson:"likes,omitempty"`
}

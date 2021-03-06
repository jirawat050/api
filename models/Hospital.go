package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Away      string        `json:"away" bson:"away"`
	Link      string        `json:"link" bson:"link"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

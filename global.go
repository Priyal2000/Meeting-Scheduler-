package main

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var lock sync.Mutex

//dskip : Stores the value of the default offset
var dskip = int64(0)

//dlimit : Stores the value of the default limit
var dlimit = int64(10)

var skip = Defaultskip
var limit = Defaultlimit

//participants : Stores the details of a participant
type participants struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Rsvp  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}

//meetings : Stores the details of a meeting
type meetings struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Participants []participants      `json:"participants,omitempty" bson:"participants,omitempty"`
	Starttime    string             `json:"starttime,omitempty" bson:"starttime,omitempty"`
	Endtime      string             `json:"endtime,omitempty" bson:"endtime,omitempty"`
	Creationtime string             `json:"creationtime,omitempty" bson:"creationtime,omitempty"`
}

func (person *participants) cons() {
	if person.Rsvp == "" {
		person.Rsvp = "Not Answered"
	}
	if person.Email == "" {
		person.Email = "defaultmail@email.com"
	}
	if person.Name == "" {
		person.Name = person.Email
	}
}

func (obj *meetings) def() {
	if obj.Title == "" {
		obj.Title = "Untitled Meeting"
	}
	if obj.Starttime == "" {
		obj.Starttime = string(time.Now().Format(time.RFC3339))
	}
	if obj.Endtime == "" {
		obj.Endtime = string(time.Now().Local().Add(time.Hour * time.Duration(1)).Format(time.RFC3339))
	}
	if obj.Creationtime == "" {
		obj.Creationtime = string(time.Now().Format(time.RFC3339))
	}
	for i := range obj.participants {
		obj.participants[i].cons()
	}
}

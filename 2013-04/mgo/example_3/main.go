package main

import (
	"fmt"
	"github.com/kr/pretty"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	Todo struct {
		Id        bson.ObjectId `bson:"_id"`
		Task      string        `bson:"t"`
		Created   time.Time     `bson:"c"`
		Updated   time.Time     `bson:"u,omitempty"`
		Completed time.Time     `bson:"cp,omitempty"`
	}
)

func main() {
	var (
		mongoSession *mgo.Session
		database     *mgo.Database
		collection   *mgo.Collection
		changeInfo   *mgo.ChangeInfo
		err          error
	)

	if mongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}

	database = mongoSession.DB("mgo_examples_03")
	collection = database.C("todos")

	// START OMIT
	var todo = Todo{
		Id:      bson.NewObjectId(),
		Task:    "Demo mgo",
		Created: time.Now(),
	}

	// This is a shortcut to collection.Upsert(bson.M{"_id": todo.id}, &todo)
	if changeInfo, err = collection.UpsertId(todo.Id, &todo); err != nil {
		panic(err)
	}
	// END OMIT

	fmt.Printf("Todo: %# v", pretty.Formatter(todo))
	fmt.Printf("Change Info: %# v", pretty.Formatter(changeInfo))

}

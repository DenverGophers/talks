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

	database = mongoSession.DB("mgo_examples_04")
	collection = database.C("todos")

	var todo = Todo{
		Id:      bson.NewObjectId(),
		Task:    "Demo mgo",
		Created: time.Now(),
	}

	// This is a shortcut to collection.Upsert(bson.M{"_id": todo.id}, &todo)
	if changeInfo, err = collection.UpsertId(todo.Id, &todo); err != nil {
		panic(err)
	}

	fmt.Printf("Todo: %# v", pretty.Formatter(todo))
	fmt.Printf("Change Info: %# v", pretty.Formatter(changeInfo))

	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"cp": time.Now(),
			}}}
	if changeInfo, err = collection.FindId(todo.Id).Apply(change, &todo); err != nil {
		panic(err)
	}

	fmt.Printf("Todo: %# v", pretty.Formatter(todo))
	fmt.Printf("Change Info: %# v", pretty.Formatter(changeInfo))
}

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
		Due       time.Time     `bson:"d,omitempty"`
		Completed time.Time     `bson:"cp,omitempty"`
	}
	TodoDueCounts struct {
		Id    int `bson:"_id"`
		Count int `bson:"count"`
	}
)

func main() {
	var (
		mongoSession *mgo.Session
		database     *mgo.Database
		collection   *mgo.Collection
		err          error
	)

	if mongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}

	database = mongoSession.DB("mgo_examples_05")
	collection = database.C("todos")

	var todos []Todo
	todos = append(todos, Todo{Id: bson.NewObjectId(), Task: "First task for today", Created: time.Now(), Due: time.Now().Add(time.Hour * 24)})
	todos = append(todos, Todo{Id: bson.NewObjectId(), Task: "Second task for today", Created: time.Now(), Due: time.Now()})
	todos = append(todos, Todo{Id: bson.NewObjectId(), Task: "Third task for today", Created: time.Now(), Due: time.Now()})
	todos = append(todos, Todo{Id: bson.NewObjectId(), Task: "Fourth task for today", Created: time.Now(), Due: time.Now()})
	todos = append(todos, Todo{Id: bson.NewObjectId(), Task: "Fifth task for today", Created: time.Now(), Due: time.Now()})

	for _, todo := range todos {
		if _, err = collection.UpsertId(todo.Id, &todo); err != nil {
			panic(err)
		}
	}
	pipeline := []bson.M{
		{"$group": bson.M{
			"_id":   bson.M{"$dayOfYear": "$d"},
			"count": bson.M{"$sum": 1},
		}},
	}

	var (
		result  TodoDueCounts
		results []TodoDueCounts
	)

	iter := collection.Pipe(pipeline).Iter()
	for {
		if iter.Next(&result) {
			results = append(results, result)
		} else {
			break
		}
	}
	err = iter.Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v", pretty.Formatter(results))
}

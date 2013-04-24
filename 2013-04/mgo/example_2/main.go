package main

import (
	"fmt"
	"github.com/kr/pretty"
	"time"
)

type (
	// START OMIT
	Todo struct {
		Task      string    `bson:"t"`
		Created   time.Time `bson:"c"`
		Updated   time.Time `bson:"u,omitempty"`
		Completed time.Time `bson:"cp,omitempty"`
	}
	// END OMIT
)

func main() {
	var todo = Todo{
		Task:    "Demo mgo",
		Created: time.Now(),
	}
	fmt.Printf("Todo: %# v", pretty.Formatter(todo))
}

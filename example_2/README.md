> Example 2
==========

Create a basic struct with annotations to allow for saving to mongoDB



	package main

	import (
		"fmt"
		"github.com/kr/pretty"
		"time"
	)

	type (
		Todo struct {
			Task      string    `bson:"t"`
			Created   time.Time `bson:"c"`
			Updated   time.Time `bson:"u,omitempty"`
			Completed time.Time `bson:"cp,omitempty"`
		}
	)

	func main() {
		var todo = Todo{
			Task:    "Demo mgo",
			Created: time.Now(),
		}
		fmt.Printf("Todo: %# v", pretty.Formatter(todo))
	}


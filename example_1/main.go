package main

import (
	"fmt"
	"github.com/kr/pretty"
	"labix.org/v2/mgo"
)

func main() {

	var (
		mongoSession *mgo.Session
		database     *mgo.Database
		collection   *mgo.Collection
		err          error
	)

	// START OMIT
	if mongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}
	database = mongoSession.DB("mgo_examples")
	collection = database.C("to_dos")
	// END OMIT

	fmt.Printf("Collection: %# v", pretty.Formatter(collection))
}

package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// START OMIT
type (
	Todos []Todo
	Todo  struct {
		Id        bson.ObjectId `json:"id"           bson:"_id"`
		Task      string        `json:"t"            bson:"t"`
		Created   time.Time     `json:"c"            bson:"c"`
		Updated   time.Time     `json:"u,omitempty"  bson:"u,omitempty"`  // NOTE: omitempty does not work on an empty time for json...
		Due       time.Time     `json:"d,omitempty"  bson:"d,omitempty"`  // Refer to this bug for details: https://code.google.com/p/go/issues/detail?id=4357
		Completed time.Time     `json:"cp,omitempty" bson:"cp,omitempty"` // Side note... time.Time is a struct...
	}
)

/* 
	NOTE: omitempty does not work on an empty time for json...                         
	Refer to this bug for details: https://code.google.com/p/go/issues/detail?id=4357  
	Side note... time.Time is a struct...                                              
*/
// END OMIT

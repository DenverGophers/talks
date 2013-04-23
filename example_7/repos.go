package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	todoRepo struct {
		Collection *mgo.Collection
	}
)

func (r todoRepo) All() (todos Todos, err error) {
	err = r.Collection.Find(bson.M{}).All(&todos)
	return
}

func (r todoRepo) Create(todo *Todo) (err error) {
	if todo.Id.Hex() == "" {
		todo.Id = bson.NewObjectId()
	}
	if todo.Created.IsZero() {
		todo.Created = time.Now()
	}
	todo.Updated = time.Now()
	_, err = r.Collection.UpsertId(todo.Id, todo)
	return
}

func (r todoRepo) Update(todo *Todo) (err error) {
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"u": time.Now(),
				"d": todo.Due,
				"t": todo.Task,
			}}}
	_, err = r.Collection.FindId(todo.Id).Apply(change, todo)

	return
}
func (r todoRepo) Destroy(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

func (r todoRepo) Complete(id string) (todo Todo, err error) {
	bid := bson.ObjectIdHex(id)
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"cp": time.Now(),
			}}}
	_, err = r.Collection.FindId(bid).Apply(change, &todo)

	return
}

func (r todoRepo) Uncomplete(id string) (todo Todo, err error) {
	bid := bson.ObjectIdHex(id)
	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$unset": bson.M{
				"cp": 1,
			}}}
	_, err = r.Collection.FindId(bid).Apply(change, &todo)

	return
}

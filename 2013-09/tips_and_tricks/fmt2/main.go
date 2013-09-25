package main

import (
	"fmt"
)

type (
	Customer struct {
		Name   string
		Street []string
		City   string
		State  string
		Zip    string
	}
	Item struct {
		Id       int
		Name     string
		Quantity int
	}
	Items []Item
	Order struct {
		Id       int
		Customer Customer
		Items    Items
	}
)

func main() {
	var items Items
	items = append(items, Item{Id: 100, Name: "Mousestrap", Quantity: 10})
	items = append(items, Item{Id: 101, Name: "USB Cable", Quantity: 2})
	order := Order{
		Id: 100,
		Customer: Customer{
			Name:   "Cory LaNou",
			Street: []string{"1062 Delaware St."},
			City:   "Denver",
			State:  "CO",
			Zip:    "80204",
		},
		Items: items,
	}

	// START OMIT
	fmt.Printf("%s\n\n", order)
	fmt.Printf("%v\n\n", order)
	fmt.Printf("%+v\n\n", order)
	fmt.Printf("%#v\n\n", order)
	fmt.Printf("%T\n", order)
	// END OMIT

}

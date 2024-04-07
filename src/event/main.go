package main

import "fmt"

func main() {
	var e Event

	// create
	e = Event{KindCreate, DataCreate{"id1", "nash"}}
	e.Print()

	// update
	e = Event{KindUpdate, DataUpdate{"id1", "lunash"}}
	e.Print()

	// delete
	e = Event{KindDelete, DataDelete{"id1"}}
	e.Print()
}

type Event struct {
	kind Kind
	data interface{}
}

type Data interface {
	DataCreate | DataDelete | DataUpdate
}

type DataDelete struct {
	ID string
}

type DataCreate struct {
	ID   string
	Name string
}

type DataUpdate struct {
	ID   string
	Name string
}

type User struct {
	ID   string
	Name string
}

type Kind string

const (
	KindCreate Kind = "add"
	KindUpdate Kind = "udate"
	KindDelete Kind = "delete"
)

func (e *Event) Print() error {
	_, error := fmt.Println("e is %+w", e)
	return error
}

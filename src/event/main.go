package main

import "fmt"

type Event struct {
	kind Kind
	// data Data
}

type Data interface {
	DataAdd | DataDelete | DataUpdate
	String() string
}

type DataDelete struct {
	ID string
}

type DataAdd struct {
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

// func (d *Data) String() string {
// 	return ""
// }

func (e *Event) Print() error {
	_, error := fmt.Println("e is %+w", e)
	return error
}

func main() {
	var e Event
	// create
	e = Event{KindCreate}
	e.Print()

	// update
	e = Event{KindUpdate}
	e.Print()

	// delete
	e = Event{KindDelete}
	e.Print()
}

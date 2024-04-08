package event

import "fmt"

type Event struct {
	Kind Kind
	Data Data
}

type Kind string

const (
	KindCreate Kind = "add"
	KindUpdate Kind = "udate"
	KindDelete Kind = "delete"
)

type Data interface {
	IsData_()
}

type DataDelete struct {
	ID string
}

func (d *DataDelete) IsData_() {}

type DataCreate struct {
	ID   string
	Name string
}

func (d *DataCreate) IsData_() {}

type DataUpdate struct {
	ID   string
	Name string
}

func (d *DataUpdate) IsData_() {}

type User struct {
	ID   string
	Name string
}

func (e *Event) Print() error {
	_, error := fmt.Println("e is %+w", e)
	return error
}

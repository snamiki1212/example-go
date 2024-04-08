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
	// guard method to constraint for only specific struct.
	_isData()
}

type DataDelete struct {
	ID string
}

func (d *DataDelete) _isData() {}

type DataCreate struct {
	ID   string
	Name string
}

func (d *DataCreate) _isData() {}

type DataUpdate struct {
	ID   string
	Name string
}

func (d *DataUpdate) _isData() {}

type User struct {
	ID   string
	Name string
}

func (e *Event) Print() error {
	_, error := fmt.Println("e is %+w", e)
	return error
}

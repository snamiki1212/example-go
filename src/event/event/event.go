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

// Print
// - interface pattern
func (e *Event) Print_InterfacePattern() {
	fmt.Println("e is %+w", e)
}

// Print
// - switch pattern
func (e *Event) Print_SwitchPattern() {
	switch d := e.Data.(type) {
	case *DataCreate:
		fmt.Println("--data-create %w / %w", d.ID, d.Name)
	case *DataUpdate:
		fmt.Println("--data-update %w / %w", d.ID, d.Name)
	case *DataDelete:
		fmt.Println("--data-delete %w", d.ID)
	}
}

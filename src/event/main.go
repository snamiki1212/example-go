package main

import (
	"event/event"
	"fmt"
)

func main() {
	var e *event.Event
	var d event.Data

	{ // interface pattern
		// create
		d = &event.DataCreate{ID: "id1", Name: "nash"}
		e = &event.Event{Kind: event.KindCreate, Data: d}
		e.Print()

		// update
		d = &event.DataUpdate{ID: "id1", Name: "lunash"}
		e = &event.Event{Kind: event.KindUpdate, Data: d}
		e.Print()

		// delete
		d = &event.DataDelete{ID: "id1"}
		e = &event.Event{Kind: event.KindDelete, Data: d}
		e.Print()
	}

	{
		// switch pattern
		switch x := d.(type) {
		case *event.DataCreate:
			x = x
			fmt.Println("--data-create")
		case *event.DataUpdate:
			fmt.Println("--data-update")
		case *event.DataDelete:
			fmt.Println("--data-delete")
		}
	}
}

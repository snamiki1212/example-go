package main

import (
	"union/event"
)

func main() {
	var e *event.Event
	var d event.Data

	// create
	d = &event.DataCreate{ID: "id1", Name: "nash"}
	e = &event.Event{Kind: event.KindCreate, Data: d}
	e.Print_InterfacePattern()
	e.Print_SwitchPattern()

	// update
	d = &event.DataUpdate{ID: "id1", Name: "lunash"}
	e = &event.Event{Kind: event.KindUpdate, Data: d}
	e.Print_InterfacePattern()
	e.Print_SwitchPattern()

	// delete
	d = &event.DataDelete{ID: "id1"}
	e = &event.Event{Kind: event.KindDelete, Data: d}
	e.Print_InterfacePattern()
	e.Print_SwitchPattern()

}

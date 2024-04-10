package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID   string `validate:"required"`
	Name string `validate:"required"`
}

func main() {
	us := &User{
		// ID:   "uuid",
		// Name: "name",
	}
	err := validator.New().Struct(us)
	if err != nil {
		fmt.Println(err)
		/**
		go run ./main.go
		Key: 'User.ID' Error:Field validation for 'ID' failed on the 'required' tag
		Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
		**/
		return
	}

	fmt.Println("OK")
}

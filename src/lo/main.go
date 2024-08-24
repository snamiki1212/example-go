package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})
	fmt.Println(names)
	us := Users{{UserID: "1", Name: "John"}, {UserID: "2", Name: "Taro"}}
	us = lo.Map(us, func(u *User, _ int) *User {
		u.Name = u.Name + "-san"
		return u
	})
	fmt.Println(us)
}

type User struct {
	UserID string
	Name   string
}

type Users []*User

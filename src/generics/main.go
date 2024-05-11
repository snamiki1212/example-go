package main

import "fmt"

type (
	User struct {
		ID  string
		Age int
	}
	Users []*User
)

func (us Users) IDs() []string {
	ids := make([]string, 0, len(us))
	for i := range us {
		ids = append(ids, us[i].ID)
	}
	return ids
}

func main() {
	fmt.Printf("hello world :D")
	us := Users{{ID: "1", Age: 11}, {ID: "2", Age: 22}, {ID: "3", Age: 33}, {ID: "4", Age: 44}}
	us = FilterBy(us, func(u *User) bool {
		return u.Age%2 == 0
	})
	fmt.Println(us.IDs())
}

func FilterBy[T any](sli []*T, fn func(x *T) bool) []*T {
	dst := make([]*T, 0, len(sli))
	for i := range sli {
		if fn(sli[i]) {
			dst = append(dst, sli[i])
		}
	}
	return dst
}

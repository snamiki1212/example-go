// functional option design pattern
package main

import (
	"fmt"
	"functional_option/user"
)

func main() {
	u, err := user.NewUser(
		user.ID("test-id-1"),
		user.FirstName("Taro"),
		user.LastName("Tanaka"),
		user.Age(18),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.FullName())
}

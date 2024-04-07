package user

import "fmt"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int64
}

type Option func(*User) error

func ID(id string) Option {
	return func(u *User) error {
		u.ID = id
		return nil
	}
}

func FirstName(fn string) Option {
	return func(u *User) error {
		if len(fn) >= 10 {
			return fmt.Errorf("FirstName err: max 10")
		}
		u.FirstName = fn
		return nil
	}
}

func LastName(ln string) Option {
	return func(u *User) error {
		if len(ln) >= 10 {
			return fmt.Errorf("FirstName err: max 10")
		}
		u.LastName = ln
		return nil
	}
}

func Age(ag int64) Option {
	return func(u *User) error {
		if ag > 150 {
			return fmt.Errorf("Age err: max 150")
		}
		u.Age = ag
		return nil
	}
}

func NewUser(setters ...Option) (*User, error) {
	args := &User{}
	for _, setter := range setters {
		if err := setter(args); err != nil {
			return nil, err
		}
	}
	return args, nil
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

//

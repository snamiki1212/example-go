package user

type User struct {
	UserID    string
	Name      string
	Age       int64
	PtrAge    *int64
	Posts     Posts
	Posts2    Posts
	callback5 func(head string, tail ...bool) (num int64)
	callback6 func(head string, tail ...func(x1 string) (y1 string)) (num int64)
}

//go:generate go run gen_slice_accessor -entity=User -slice=Users -exclude=Posts,Posts2 -in=user.go -out=users_accessor_gen.go
type Users []*User

type Post struct {
	PostID string
}

type Posts []*Post

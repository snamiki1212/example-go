package user

type User struct {
	UserID string
	Name   string
	Age    int64
	PtrAge *int64
	Posts  Posts
}

//go:generate go run slice_accessor -entity=User -slice=Users
type Users []*User

type Post struct {
	PostID string
}

type Posts []*Post

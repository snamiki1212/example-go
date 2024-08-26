package user

type User struct {
	UserID string
	Name   string
	Age    int64
	PtrAge *int64
	Posts  Posts
	Posts2 Posts
}

//go:generate go run slice_accessor -entity=User -slice=Users -exclude=Posts,Posts2
type Users []*User

type Post struct {
	PostID string
}

type Posts []*Post

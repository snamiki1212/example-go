package user

type User struct {
	UserID string
	Name   string
	Age    int64
	PtrAge *int64
	Posts  Posts
}

//go:generate go run slicer -entity=User -slices=Users
type Users []*User

type Post struct {
	PostID string
}

type Posts []*Post

// UserIDs
func (us Users) UserIDs() []string {
	ids := make([]string, 0, len(us))
	for i := range us {
		ids = append(ids, us[i].UserID)
	}
	return ids
}

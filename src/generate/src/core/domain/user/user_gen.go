// Code generated by go generate DO NOT EDIT.

package user

// UserIDs
func (xs Users) UserIDs() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].UserID)
	}
	return sli
}

// Names
func (xs Users) Names() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Name)
	}
	return sli
}

// Ages
func (xs Users) Ages() []int64 {
	sli := make([]int64, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Age)
	}
	return sli
}

// PtrAges
func (xs Users) PtrAges() []*int64 {
	sli := make([]*int64, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].PtrAge)
	}
	return sli
}

// Posts2s
func (xs Users) Posts2s() []Posts {
	sli := make([]Posts, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Posts2)
	}
	return sli
}


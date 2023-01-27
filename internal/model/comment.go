package model

type Comment struct {
	ID       int64
	PostId   int64
	Username string
	Message  string
	Like     int
	Dislike  int
	Born     string
}

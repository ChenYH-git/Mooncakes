package models

type User struct {
	UserID   int64  `db:"user_id"`
	Score    int64  `db:"score"`
	Chances  int8   `db:"chances"`
	Username string `db:"username"`
	Password string `db:"password"`
	Gift     string `db:"gift"`
	Token    string
}

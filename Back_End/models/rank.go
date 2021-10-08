package models

type Me struct {
	Username string `form:"username" db:"username"`
	Score    int64  `form:"score" db:"score"`
	Chances  int8   `form:"chances" db:"chances"`
}

type ApiRankDetail struct {
	Username string `form:"username" db:"username"`
	Gift     string `form:"gift" db:"gift"`
	Score    int64  `form:"score" db:"score"`
}

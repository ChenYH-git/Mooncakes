package models

type DiceRes struct {
	Score int64 `form:"score"`
}

type Chance struct {
	Chances int8 `db:"chances"`
}

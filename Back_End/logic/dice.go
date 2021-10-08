package logic

import (
	"Mooncakes/dao/mysql"
	"Mooncakes/models"
)

func AddScore(userID int64, res *models.DiceRes) (err error) {
	if err = mysql.AddScore(userID, res); err != nil {
		return err
	}
	if err = mysql.DelChance(userID); err != nil {
		return err
	}
	return
}

func CheckChances(userID int64) (bool, error) {
	data, err := mysql.CheckChances(userID)
	if err != nil {
		return false, err
	}
	if data[0].Chances <= 0 {
		return true, err
	}
	return false, err
}

package mysql

import (
	"Mooncakes/models"
	"Mooncakes/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(settings.Conf.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(settings.Conf.MySQLConfig.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}

func GetRankList() (data []*models.ApiRankDetail, err error) {
	sqlStr := `select 
	username, gift, score
	from user
	order by score
	desc
	`
	data = make([]*models.ApiRankDetail, 0, 15)
	err = db.Select(&data, sqlStr)
	return
}

func GetMy(userID int64) (data []*models.Me, err error) {
	sqlStr := `select username, score, chances from user where user_id = ?`
	data = make([]*models.Me, 0, 2)
	err = db.Select(&data, sqlStr, userID)
	return data, err
}

func AddScore(userID int64, res *models.DiceRes) (err error) {
	sqlStr := `update user set score = score + ? where user_id = ?`
	_, err = db.Exec(sqlStr, res.Score, userID)
	return err
}

func DelChance(userID int64) (err error) {
	sqlStr := `update user set chances = chances - 1 where user_id = ?`
	_, err = db.Exec(sqlStr, userID)
	return err
}

func CheckChances(userID int64) (data []*models.Chance, err error) {
	sqlStr := `select chances from user where user_id = ?`
	data = make([]*models.Chance, 0, 2)
	err = db.Select(&data, sqlStr, userID)
	return data, err
}

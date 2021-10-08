package mysql

import (
	"Mooncakes/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "jiami"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)

	sqlStr := `insert into user(user_id, username, password, gift, score, chances) values(?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Gift, user.Score, user.Chances)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}

	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorUserPwd
	}
	return
}

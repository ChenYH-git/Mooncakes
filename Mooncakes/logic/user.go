package logic

import (
	"Mooncakes/dao/mysql"
	"Mooncakes/models"
	"Mooncakes/pkg/jwt"
	"Mooncakes/pkg/snowflakes"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	userID := snowflakes.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Gift:     "none",
		Score:    0,
		Chances:  3,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.Login(user); err != nil {
		return
	}

	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

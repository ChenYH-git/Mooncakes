package logic

import (
	"Mooncakes/dao/mysql"
	"Mooncakes/models"
)

const Num = 15

var gift = [Num]string{"漫步者耳机", "IKBC W200键盘", "网易严选按摩椅", "Keep年卡",
	"膳魔师保温杯", "小米屏幕挂灯", "罗技K380无线键盘", "紫米便携充电宝", "原味鸡x10",
	"紫米快速充电头", "闪迪双接口U盘", "零食礼包", "洗衣液", "纸巾", "口罩",
}

func GetRank() (data []*models.ApiRankDetail, err error) {
	data, err = mysql.GetRankList()
	if err != nil {
		return nil, err
	}
	for i, v := range data {
		if i > Num-1 {
			break
		}
		v.Gift = gift[i]
	}
	return
}

func GetMyRank(userID int64) (data []*models.Me, err error) {
	data, err = mysql.GetMy(userID)
	if err != nil {
		return nil, err
	}
	return
}

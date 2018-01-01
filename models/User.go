package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Email    string `orm:"size(255)"`
	RealName string `orm:"size(255)"`
	Password string `orm:"size(255)"`
	IsSuper  bool
	Status   int8
	Imgsrc   string `orm:"size(255)"`
}

// 设置引擎为 MyISAM
func (self *User) TableEngine() string {
	return "MyISAM"
}

//插入一个用户，成功返回id,失败返回0
func InsertUserInfo(user *User) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err == nil {
		return id
	}
	return 0
}

func GetUserList() []User {
	o := orm.NewOrm()
	var list []User
	_, err := o.Raw("select * from t_user").QueryRows(&list)
	if err == nil {
		return list
	}
	return nil
}

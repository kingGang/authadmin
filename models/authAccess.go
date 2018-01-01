package models

type AuthGroupAccess struct {
	Id      int
	Uid     int
	GroupId int
}

func (self *AuthGroupAccess) TableEngine() string {
	return "MyISAM"
}

// 多字段唯一键
func (self *AuthGroupAccess) TableUnique() [][]string {
	return [][]string{
		[]string{"Uid", "GroupId"},
	}
}

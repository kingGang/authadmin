package models

type AuthGroup struct {
	Id     int
	Title  string `orm:"size(100)"`
	Status int8   `orm:default(1)`
	Rules  string `orm:"size(80)"`
}

func (self *AuthGroup) TableEngine() string {
	return "MyISAM"
}

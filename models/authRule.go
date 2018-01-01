package models

//-- id:主键，name：规则唯一标识, title：规则中文名称 status 状态：为1正常，为0禁用，condition：规则表达式，为空表示存在就验证，不为空表示按照条件验证
type AuthRule struct {
	//权限ID
	Id        int
	Name      string `orm:"size(80);default('')"`
	Title     string `orm:"size(20);default('')"`
	Type      int8   `orm:"default(1)"`
	Status    int8   `orm:"default(1)"`
	Condition string `orm:"null;size(100)"`
}

func (m *AuthRule) TableEngine() string {
	return "MyISAM"
}

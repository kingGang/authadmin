package common

import (
	"github.com/astaxie/beego/orm"
	"lingo8.cn/authadmin/models"
	//引用MySQL开发包
	_ "github.com/go-sql-driver/mysql"
)

//注册MySQL驱动，设置链接数据库
func Initdb() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/rbac1?charset=utf8")
	orm.RegisterModelWithPrefix("t_", new(models.User), new(models.AuthGroup), new(models.AuthRule), new(models.AuthGroupAccess))
	orm.RunSyncdb("default", false, true)
}

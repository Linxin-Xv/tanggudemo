package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id       int
	Name     string
	Password string
}

// 表的设计有关内容
func init() {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/tanggu?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(User))
	//生成表
	orm.RunSyncdb("default", false, true)
}

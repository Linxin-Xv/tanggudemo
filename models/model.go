package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id        int    `json:"user_id"`
	Name      string `orm:"size(32);unique" json:"name"`
	Password  string `orm:"size(128)"json:"password"`
	Mobile    string `orm:"size(11)" json:"mobile"`
	Id_card   string `orm:"size(20)" json:"id_card"`
	Type      string `orm:"size(128)" json:"type"`
	Level     int    `json:"level"`
	Vip_level int    ` json:"viplevel"`
}
type Email struct {
	Id      int    `json:"email_id"`
	Address string `orm:"size(128)" json:"email_address"`
}

// 表的设计有关内容
func init() {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/tanggu?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(User), new(Email))
	//生成表
	orm.RunSyncdb("default", false, true)
}

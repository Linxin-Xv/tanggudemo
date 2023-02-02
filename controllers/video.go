package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "tangguwebdemo/models"
)

type VideoController struct {
	beego.Controller
}

func (c *VideoController) Get() {
	////1.有orm对象
	//o := orm.NewOrm()
	////2.有一个要插入数据的结构体对象
	//user := models.User{}
	////3.对结构体对象赋值
	//user.Name = "xlx"
	//user.Password = "030503"
	////4.插入
	//_, err := o.Insert(&user)
	//if err != nil {
	//	println(err)
	//	return
	//}
	//user.Id = 1
	//err := o.Read(&user)
	//if err != nil {
	//	println(err)
	//	return
	//}
	//
	//user.Id = 1
	//err := o.Read(&user)
	//if err == nil {
	//	user.Name = "徐林鑫"
	//}
	//_, err = o.Update(&user)
	//if err != nil {
	//	println(err)
	//}
	//user.Id = 1
	//_, err := o.Delete(&user)
	//if err == nil {
	//	println("删除成功！")
	//}
	//c.Data["Website"] = "tanggu.cn"
	//c.Data["Email"] = "xlxmatrix@qq.com"
	//c.Data["data"] = "欢迎来到汤谷汉教平台！"
	//c.TplName = "index.html"
}

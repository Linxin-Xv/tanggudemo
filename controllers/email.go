package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"tangguwebdemo/models"
	_ "tangguwebdemo/models"
)

type EmailController struct {
	beego.Controller
}

func (c *EmailController) HandleEmail() {
	orm.Debug = true
	//1.拿到数据
	Email := c.GetString("email") //因为是从前端拿数据，这个名称需要跟前端中的名称保持一致
	//2.对数据进行校验
	if Email == "" {
		println("数据不能为空")
		c.Redirect("/", 302) //重定向函数，如果发生错误页面重新回到/register，并返回错误码302
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	email := models.Email{}
	email.Address = Email
	_, err := o.Insert(&email)
	if err != nil {
		println("插入数据库失败")
		c.Redirect("/", 302)
		return
	}
	//4.返回登陆界面
	c.TplName = "index.html" //指定视图文件，同时可以给这个视图传递一些数据如在c.Data["errmsg"]，优点就是能够传递数据
	c.Redirect("/", 302)     //跳转，不能传递数据。优点是速度快
}

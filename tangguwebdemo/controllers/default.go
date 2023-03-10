package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"tangguwebdemo/models"
	_ "tangguwebdemo/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	////1.有orm对象
	//o := orm.NewOrm()
	////2.有一个要插入数据的结构体对象
	//user := models.User{}
	////3.对结构体对象赋值
	//user.Name = "田静宜"
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
	//	user.Name = "可爱的小田"
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
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	//1.拿到数据
	userName := c.GetString("userName") //因为是从前端拿数据，这个名称需要跟前端中的名称保持一致
	password := c.GetString("password") //因为是从前端拿数据，这个名称需要跟前端中的名称保持一致
	//2.对数据进行校验
	if userName == "" || password == "" {
		println("数据不能为空")
		c.Redirect("/register", 302) //重定向函数，如果发生错误页面重新回到/register，并返回错误码302
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Password = password
	_, err := o.Insert(&user)
	if err != nil {
		println("插入数据库失败")
		c.Redirect("/register", 302)
		return
	}
	//4.返回登陆界面
	c.TplName = "login.html"  //指定视图文件，同时可以给这个视图传递一些数据如在c.Data["errmsg"]，优点就是能够传递数据
	c.Redirect("/login", 302) //跳转，不能传递数据。优点是速度快
}
func (c *MainController) ShowRegister() {
	c.TplName = "register.html"
}
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}
func (c *MainController) HandleLogin() {
	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.判断数据是否合法
	if userName == "" || pwd == "" {
		println("输入数据不合法")
		c.TplName = "login.html"
		return
	}
	//3.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		println("查询失败")
		c.TplName = "login.html"
		c.Ctx.WriteString("用户名或密码错误，登陆失败！")
		return
	}
	user.Password = pwd
	err = o.Read(&user, "Password")
	if err != nil {
		println("查询失败")
		c.TplName = "login.html"
		c.Ctx.WriteString("用户名或密码错误，登陆失败！")
		return
	}
	//4.跳转
	c.Ctx.WriteString("登陆成功，欢迎您")
}

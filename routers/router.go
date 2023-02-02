package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tangguwebdemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainregisterController{})
	beego.Router("/register", &controllers.MainregisterController{}, "post:HandleRegister")
	//注意：当实现了自定义的get请求方法，请求将不会访问默认方法
	beego.Router("/login", &controllers.MainregisterController{}, "post:HandleLogin")
	beego.Router("/email", &controllers.EmailController{}, "post:HandleEmail")
}

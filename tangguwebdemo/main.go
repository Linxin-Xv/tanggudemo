package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "tangguwebdemo/models"
	_ "tangguwebdemo/routers"
)

func main() {
	beego.Run()
}

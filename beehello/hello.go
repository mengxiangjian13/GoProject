package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString("hello,world!")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

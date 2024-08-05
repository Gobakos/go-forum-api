package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {

	c.TplName = "index.html"
	if err := c.Render(); err != nil {
		c.Ctx.WriteString("Error rendering template: " + err.Error())
	}
}

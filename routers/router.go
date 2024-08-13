// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
    "api/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

/* Always care for HTTP Parameter Pollution,when you have a front-end (reverse proxy) */
/* In JSON format,the application interpretates 2nd occurrence */



func init() {

    beego.InsertFilter("/v1/api/posts/:post_id", beego.BeforeRouter, middleware.AuthMiddleware)
    beego.InsertFilter("/v1/api/users/:id", beego.BeforeRouter, middleware.AuthMiddleware)
		ns := beego.NewNamespace("/v1/api",
			beego.NSNamespace("/users",
				beego.NSRouter("/", &controllers.UserController{}, "get:GetUsers"),
				beego.NSRouter("/:id", &controllers.UserController{}, "get:GetUser"),
				beego.NSRouter("/:id", &controllers.UserController{}, "delete:Delete"),
			),
			beego.NSNamespace("/auth",
				beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
				beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
			),
			beego.NSNamespace("/posts",
			beego.NSRouter("/", &controllers.PostController{}, "post:AddPost"),
			beego.NSRouter("/:post_id", &controllers.PostController{}, "delete:DeletePost"),
			),
		)
	
		// Add the namespace to the Beego router
		beego.AddNamespace(ns)
		beego.Router("/", &controllers.MainController{})
	}
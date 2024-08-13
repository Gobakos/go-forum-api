package controllers

import (
	"api/services"
	//"test/models"
	//"strings"
	//"encoding/json"
	//"github.com/golang-jwt/jwt/v5"
	//"os"
	// "test/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type PostController struct {
	beego.Controller
}

func (p *PostController) AddPost() {
	// var claims *controllers.claims
	JWT:=p.Ctx.Input.Header("Authorization")
	err := services.IsValidPost(p.Ctx.Input.RequestBody,JWT)
	if err != nil {
		p.Data["json"] = map[string]string{"Error": err.Error()}
		p.Ctx.Output.SetStatus(400)
		p.ServeJSON()
		return
	}
	p.Data["json"] = map[string]string{"message": "Post added successfully"}
	p.Ctx.Output.SetStatus(201)
	p.ServeJSON()
}

func (p *PostController) DeletePost() {
    postID := p.GetString(":post_id")
    if postID == "" {
        p.Data["json"] = map[string]string{"error": "Post ID is required"}
        p.Ctx.Output.SetStatus(400)
        p.ServeJSON()
        return
    }
    
    err := services.DeletePost(postID)
    if err != nil {
        p.Data["json"] = map[string]string{"error": err.Error()}
        p.Ctx.Output.SetStatus(400)
        p.ServeJSON()
        return
    }

    p.Data["json"] = map[string]string{"message": "Post deleted successfully"}
    p.Ctx.Output.SetStatus(200)
    p.ServeJSON()
}
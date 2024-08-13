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
type UserController struct {
	beego.Controller
}

type ReturnUsers struct {
	Name     string `json:"name"`
    Lastname string `json:"lastname"`
	Description string `json:"description"`
}

/*			GET ALL USERS 			*/

func (u *UserController) GetUsers() {
    users, err := services.GetUsers()
    if err != nil {
        u.Data["json"] = map[string]string{"error": "Failed to fetch users"}
        u.Ctx.Output.SetStatus(500)
        u.ServeJSON()
        return
    }
    var returnUsers []ReturnUsers
    for _, user := range users {
        returnUser := ReturnUsers{
            Name:     user.Name,
            Lastname: user.Lastname,
			Description: user.Description,
        }
        returnUsers = append(returnUsers, returnUser)
    }
    u.Data["json"] = returnUsers
    u.Ctx.Output.SetStatus(200)
    u.ServeJSON()
}

/*			DELETE A USER 		*/

func (u *UserController) Delete() {
	id := u.Ctx.Input.Param(":id")
	if id == "" {
		u.Data["json"] = map[string]string{"error": "Invalid Request"}
		u.Ctx.Output.SetStatus(400)
		u.ServeJSON()
		return
	}
	deleted, err := services.DeleteUser(id)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Failed to delete user"}
		u.Ctx.Output.SetStatus(500)
		u.ServeJSON()
		return
	}
	if !deleted {
		u.Data["json"] = map[string]string{"error": "User not found"}
		u.Ctx.Output.SetStatus(404)
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]string{"message": "User deleted successfully"}
	u.Ctx.Output.SetStatus(200)
	u.ServeJSON()
}

/*			GET INFO ABOUT A USER			*/

func (u *UserController) GetUser() {
	
	id := u.Ctx.Input.Param(":id")
	if id==""{
		u.Data["json"] = map[string]string{"error": "Invalid_Request"}
		u.Ctx.Output.SetStatus(404)
		u.ServeJSON()
        return
	}

	user, err := services.GetUser(id)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Failed_to_fetch_user"}
		u.Ctx.Output.SetStatus(500)
        u.ServeJSON()
        return
	} else if user.Id == "" {
		u.Data["json"] = map[string]string{"Err": "User_cannot_found"}
		u.Ctx.Output.SetStatus(404)
        u.ServeJSON()
        return
	} else {
		u.Data["json"] = user
		u.Ctx.Output.SetStatus(200)
        u.ServeJSON()
        return
	}

}


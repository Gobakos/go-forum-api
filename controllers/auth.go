package controllers

import (
	//"test/models"
	//"encoding/json"
	// "crypto/hmac"
	// "crypto/sha256"
	"api/services"
	"strings"
	//"os"
	"time"
	//"encoding/hex"
	//b64 "encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	//"github.com/joho/godotenv"
	/*  sDec, _ := b64.StdEncoding.DecodeString(sEnc) */
	/* uEnc := b64.URLEncoding.EncodeToString([]byte(data)) */
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Auth
type AuthController struct {
	beego.Controller
}

type Claims struct {
	Username string `json:"username"`
	Id string `json:"Id"`
	Is_Admin   string   `json:"is_admin"`
	jwt.RegisteredClaims //Embedded standard
}

/* 				LOGIN				*/

func (this *AuthController) Login(){

	//Unmarshal accepts a memory location,to point there and store the data from the body;
	//If a field is unpresent,Unmarshal stores it as ""
	login, err := services.IsValidLogin(this.Ctx.Input.RequestBody)
	if err != nil {
		this.Data["json"] = map[string]string{"Error": err.Error()}
		this.Ctx.Output.SetStatus(400)
		this.ServeJSON()
		return
	}

	hashed_pass, errHash := services.HashedPasswordAuth(login.Password)

	if errHash != nil{
		this.Data["json"]=map[string]string{"Error":"Server error"}
		this.Ctx.Output.SetStatus(500)
		this.ServeJSON()
		return
	}

	boolean,id,isAdmin,err := services.TryLogin(login.Username,hashed_pass)

	if err != nil{
		this.Data["json"]=map[string]string{"Error":"Something went wrong"}
		this.Ctx.Output.SetStatus(400)
		this.ServeJSON()
		return
	}else if !boolean{
		this.Data["json"]=map[string]string{"Error":"Invalid creds"}
		this.Ctx.Output.SetStatus(403)
		this.ServeJSON()
		return
	}else{
		claims := Claims{
			Username: login.Username,
			Id: id,
			Is_Admin: isAdmin,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		}

		tokenString,err:=services.TokenGen(claims)
		if err != nil {
			this.Data["json"] = map[string]string{"Message": err.Error()}
			this.Ctx.Output.SetStatus(500)
			this.ServeJSON()
			return
		}
		this.Ctx.Output.Header("Authorization", "Bearer "+tokenString)
		this.Data["json"]=map[string]string{"Message":"Creds are valid!"}
		this.Ctx.Output.SetStatus(200)
		this.ServeJSON()
		return
	}
}


func (this *AuthController) Register(){
	register, err := services.IsValidRegister(this.Ctx.Input.RequestBody)
	if err != nil {
		this.Data["json"] = map[string]string{"Error": err.Error()}
		this.Ctx.Output.SetStatus(400)
		this.ServeJSON()
		return
	}
		//For now,we do not check the true or false!
		err=services.AddUser(register)
		if err != nil{	
			if strings.Contains(err.Error(), "1062") {
				this.Data["json"] = map[string]string{"Error": "Username already exists!"}
				this.Ctx.Output.SetStatus(400)
				this.ServeJSON()
				return
			}
			this.Data["json"]=map[string]string{"Error":err.Error()}
			this.Ctx.Output.SetStatus(500)
			this.ServeJSON()
			return
		}
		this.Data["json"]=map[string]string{"Success":"You succesfully registered!"}
		this.Ctx.Output.SetStatus(201)
		this.ServeJSON()
}
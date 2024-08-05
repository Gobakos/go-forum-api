package models

import (
	// "crypto/hmac"
	// "crypto/sha256"
	// "encoding/hex"
	//"fmt"
	"github.com/beego/beego/v2/client/orm"
	//"strconv"
	//"time"
)

func init() {
	//register the model
	orm.RegisterModel(new(Users))
}
/* Creates table named "users" with these fields */

type Users struct {
	Id       string `orm:"pk;size(50);omitempty"`
	Name string `orm:"size(100)"`
	Lastname string `orm:"size(100)"`
	Username string `orm:"size(50);unique"`
	Password string `orm:"size(100)"`
	Description string `orm:"size(150)"`
	Is_Admin string `orm:"size(25);omitempty`
	// Posts    []*Post `orm:"reverse(many)"`
}


/*             GET USERS              */

func GetUsers()([]*Users,error){
	o:=orm.NewOrm()
	var users []*Users
	_,err :=o.QueryTable("users").All(&users)
	if err != nil{
		return nil,err
	}
	return users,nil
}

/*             GET USER              */

func GetUser(id string)(Users,error){
	o:=orm.NewOrm()
	var user Users
	_,err := o.QueryTable("users").Filter("Id",id).All(&user)
	if err !=nil{
		return user,err
	}
	return user,err
}


/*             DELETE USER              */

func DelUser(id string)(bool,error){
	o:=orm.NewOrm()
	num,err:=o.QueryTable("users").Filter("Id",id).Delete()
	if err != nil {
		return false,err
	}
	if num == 0 {
        return false, nil
    }
    return true, nil
}
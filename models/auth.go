package models

import (
	"github.com/beego/beego/v2/client/orm"
	// "errors"
	"strconv"
	"time"
)

type Login struct {
	Username string `orm:"size(50)"`
	Password string `orm:"size(100)"`
}

func TryLogin(username string, password string) (bool,string, string, error) {
    o := orm.NewOrm()
    var user Users
    err := o.QueryTable("users").Filter("Username", username).Filter("Password", password).One(&user)
    var isAdmin="false"
    if err != nil {
        if err == orm.ErrNoRows {
            return false, "",isAdmin, nil
        }
        return false, "",isAdmin, err
    }
    if username=="admin"{
        isAdmin="true"
    }
    return true, user.Id,isAdmin, nil
}

/*             ADD USER              */

func AddUser(user *Users) (bool, error) {
	o := orm.NewOrm()
	user.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10) //Generate a unique ID
	user.Is_Admin="false"
	if user.Username=="admin"{
        user.Is_Admin="true"
    }
	_, err := o.Insert(user) // ORM determines which table to insert into
	if err != nil {
		return false, err
	}
	return true, nil
}
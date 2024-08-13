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
	orm.RegisterModel(new(Posts))
}

type Posts struct {
    Id          string `orm:"pk;size(50);omitempty"`
    Title       string `orm:"size(100)"`
    Username    string `orm:"size(50)"`
    Description string `orm:"size(150)"`
    User        *Users `orm:"rel(fk)"`
}

func AddPost(post *Posts) (bool, error) {
	o := orm.NewOrm()
	_, err := o.Insert(post)
	if err != nil {
		return false, err
	}
	return true, nil
}
func GetPostUserId(postID string) (string, error) {
	o := orm.NewOrm()
	var post Posts
	err := o.QueryTable("posts").Filter("id", postID).One(&post)
	if err != nil {
		return "", err
	}
	return post.User.Id, nil
}

func DeletePost(postID string) (bool, error) {
    o := orm.NewOrm()
    num, err := o.Delete(&Posts{Id: postID})
    if err != nil {
        return false, err
    }
    if num == 0 {
        return false, nil // No rows affected
    }
    return true, nil
}
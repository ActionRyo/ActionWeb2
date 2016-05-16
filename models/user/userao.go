package user

import (
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
)

type UserAoModel struct {
	web.Model
	UserDb UserDbModel
}

func (this *UserAoModel) Login(name string, password string) {
	if name == "" || password == "" {
		Throw(1, "请输入账户或密码")
	}

	users := this.UserDb.GetByNameAndPassword(name, password)
	if len(users) == 0 {
		Throw(1, "用户或密码错误")
	}

	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	sess.Set("UserId", users[0].UserId)
	defer sess.SessionRelease()
}

func (this *UserAoModel) Logout() {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}
	sess.Delete("UserId")
	defer sess.SessionRelease()
}

func (this *UserAoModel) Register(name string, password string) {
	if name == "" || password == "" {
		Throw(1, "请输入账户或密码")
	}

	users := this.UserDb.GetByName(name)
	if len(users) != 0 {
		Throw(1, "存在重复的用户名")
	}

	this.UserDb.Add(User{
		UserAccount: name,
		UserPwd:     password,
	})

	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}
	sess.Set("UserId", users[0].UserId)
	defer sess.SessionRelease()
}
package user

import (
	"github.com/fishedee/web"
)

type UserDbModel struct {
	web.Model
}

func (this *UserDbModel) GetByNameAndPassword(name string, password string) []User {
	result := []User{}
	err := this.DB.Where("userAccount = ?", name).Where("userPwd = ?", password).Find(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func (this *UserDbModel) GetByName(name string) []User {
	result := []User{}
	err := this.DB.Where("userAccount = ?", name).Find(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func (this *UserDbModel) Add(data User) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

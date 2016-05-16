package controllers

import (
	"ActionWeb2/models/book"
	"ActionWeb2/models/user"
	"fmt"
	"strconv"
)

type MainController struct {
	BaseController
	UserAo user.UserAoModel
	BookAo book.BookAoModel
}

// 主页
func (this *MainController) Main_HtmlMain() interface{} {
	books := this.BookAo.QueryBook()
	return books
}

// 登录
func (this *MainController) Login_HtmlLogin() interface{} {
	if this.Ctx.GetMethod() == "POST" {
		var input struct {
			Account  string
			Password string
		}
		this.CheckPost(&input)

		this.UserAo.Login(input.Account, input.Password)

		return redirectOut{url: "/home/main"}
	}
	return nil
}

// 注册
func (this *MainController) Regist_HtmlRegist() interface{} {
	if this.Ctx.GetMethod() == "POST" {
		var input struct {
			Account  string
			Password string
		}
		this.CheckPost(&input)

		this.UserAo.Register(input.Account, input.Password)

		return redirectOut{url: "/home/main"}
	}
	return nil
}

// 注销
func (this *MainController) Logout_HtmlLogout() interface{} {
	this.UserAo.Logout()
	return redirectOut{url: "/home/login"}
}

// 添加书本信息
func (this *MainController) AddBook_HtmlAddBook() interface{} {

	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	defer sess.SessionRelease()
	userid := sess.Get("UserId")

	if userid == nil {
		return redirectOut{url: "/home/login"}
	}

	// 转换为整形
	clientIdString := fmt.Sprintf("%v", userid)
	clientIdInt, errs := strconv.Atoi(clientIdString)
	if errs != nil {
		panic(errs)
	}

	if this.Ctx.GetMethod() == "POST" {
		var input struct {
			Code string
			Name string
		}
		this.CheckPost(&input)

		this.BookAo.AddBook(input.Code, input.Name, clientIdInt)

		return redirectOut{url: "/home/main"}
	}
	return nil
}

// 编辑图书信息
func (this *MainController) EditBook_HtmlEditBook() interface{} {

	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	userid := sess.Get("UserId")
	defer sess.SessionRelease()
	if userid == nil {
		return redirectOut{url: "/home/login"}
	}

	// 转换为整形
	clientIdString := fmt.Sprintf("%v", userid)
	clientIdInt, errs := strconv.Atoi(clientIdString)
	if errs != nil {
		panic(errs)
	}

	if this.Ctx.GetMethod() == "GET" {
		bid := this.Ctx.GetParam("bid")
		clientIdString := fmt.Sprintf("%v", bid)
		clientIdInt, err := strconv.Atoi(clientIdString)

		if err != nil {
			panic(err)
		}
		book := this.BookAo.QueryBookByID(clientIdInt)

		return book[0]
	}

	if this.Ctx.GetMethod() == "POST" {
		var input struct {
			Bookid int
			Code   string
			Name   string
		}
		this.CheckPost(&input)

		if clientIdInt == 0 {
			return redirectOut{url: "/home/login"}
		}

		this.BookAo.EditBook(input.Bookid, input.Code, input.Name, clientIdInt)

		return redirectOut{url: "/home/main"}
	}

	return nil
}

// 删除图书信息
func (this *MainController) DelBook_HtmlDelBook() interface{} {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	userid := sess.Get("UserId")
	defer sess.SessionRelease()
	if userid == nil {
		return redirectOut{url: "/home/login"}
	}

	// 获取URL的值
	bid := this.Ctx.GetParam("bid")
	clientIdString := fmt.Sprintf("%v", bid)
	clientIdInt, err := strconv.Atoi(clientIdString)
	if err != nil {
		panic(err)
	}

	this.BookAo.DeleteBook(clientIdInt)

	return redirectOut{url: "/home/main"}
}

package book

import (
	"fmt"
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
)

type BookAoModel struct {
	web.Model
	BookDb BookDbModel
}

// 添加书本
func (this *BookAoModel) AddBook(bookCode string, bookName string, userid int) {

	if bookCode == "" || bookName == "" {
		Throw(1, "图书编码或者图书名字为空")
	}

	// 判断图书编码是否存在
	books := this.BookDb.QueryBookByCodeAddMode(bookCode)
	if len(books) != 0 {
		Throw(1, "存在相同的图书编码")
	}

	this.BookDb.AddBook(Book{
		BookCode: bookCode,
		BookName: bookName,
		UserId:   userid,
	})
}

// 编辑图书
func (this *BookAoModel) EditBook(bookid int, bookcode string, bookname string, userid int) {
	if bookid == 0 {
		Throw(1, "URL参数有误")
	}

	if bookcode == "" || bookname == "" {
		Throw(1, "图书编码或者图书名字为空")
	}

	// 判断图书编码是否存在
	books := this.BookDb.QueryBookByCodeEditMode(bookcode, bookid)
	if len(books) != 0 {
		Throw(1, "存在相同的图书编码")
	}

	this.BookDb.EditBook(Book{
		BookId:   bookid,
		BookCode: bookcode,
		BookName: bookname,
		UserId:   userid,
	})
}

// 删除书本信息
func (this *BookAoModel) DeleteBook(bid int) {

	// 验证书本信息是否存在
	this.BookDb.QueryBookByID(bid)
	//if len(books) == 0 {
	//	msg := "不存在ID为：[" + string(bid) + "] 的书本信息记录"
	//	Throw(1, msg)
	//}

	// 删除书本信息
	this.BookDb.DeleteBook(bid)
}

// 根据书本ID获取书本信息
func (this *BookAoModel) QueryBookByID(bid int) []Book {
	var books = []Book{}

	// 验证书本信息是否存在
	books = this.BookDb.QueryBookByID(bid)
	clientIdString := fmt.Sprintf("%v", bid)
	if len(books) == 0 {
		msg := "不存在ID为：[" + clientIdString + "] 的书本信息记录"
		Throw(1, msg)
	}

	return books
}

// 获取所有的书本信息
func (this *BookAoModel) QueryBook() []Book {
	var books = []Book{}
	books = this.BookDb.QueryAllBook()

	return books
}

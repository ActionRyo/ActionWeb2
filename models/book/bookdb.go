package book

import (
	"github.com/fishedee/web"
)

type BookDbModel struct {
	web.Model
}

// 获取所有书本信息
func (this *BookDbModel) QueryAllBook() []Book {
	result := []Book{}
	err := this.DB.Find(&result)
	if err != nil {
		panic(err)
	}

	return result
}

// 根据书本ID获取书本信息
func (this *BookDbModel) QueryBookByID(bid int) []Book {
	result := []Book{}
	err := this.DB.Where("bookid = ?", bid).Find(&result)
	if err != nil {
		panic(err)
	}

	return result
}

// 根据图书编号获取书本信息(新增时候用)
func (this *BookDbModel) QueryBookByCodeAddMode(code string) []Book {
	result := []Book{}
	err := this.DB.Where("bookcode = ?", code).Find(&result)
	if err != nil {
		panic(err)
	}

	return result
}

// 根据图书编号获取书本信息(修改时候用)
func (this *BookDbModel) QueryBookByCodeEditMode(code string, bid int) []Book {
	result := []Book{}
	err := this.DB.Where("bookcode=?", code).Where("bookid<>?", bid).Find(&result)
	if err != nil {
		panic(err)
	}

	return result
}

// 新增书本信息
func (this *BookDbModel) AddBook(data Book) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

// 修改书本信息
func (this *BookDbModel) EditBook(data Book) {
	_, err := this.DB.Where("bookid=?", data.BookId).Update(data)

	if err != nil {
		panic(err)
	}
}

// 删除书本信息
func (this *BookDbModel) DeleteBook(bid int) {
	_, err := this.DB.Where("bookid=?", bid).Delete(&Book{})

	if err != nil {
		panic(err)
	}
}

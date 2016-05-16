package book

type Book struct {
	BookId   int
	BookCode string
	BookName string
	UserId   int
}

type ListBooks struct {
	ArrBooks []*Book
}

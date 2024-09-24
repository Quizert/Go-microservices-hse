package book

type UserBook struct {
	title  string
	author string
}

func (ub UserBook) GetTitle() string {
	return ub.title
}

func (ub UserBook) GetAuthor() string {
	return ub.author
}

func CreateUserBook(title string, author string) *UserBook {
	return &UserBook{title: title, author: author}
}

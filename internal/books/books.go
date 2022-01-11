package books

type Repository interface {
	CreateBook(book Book)
	UpdateBookByID(book Book)
	GetAllBooks() []Book
	GetByID(book Book) Book
	DeleteByID(book Book)
}

func New() Repository {
	return &repository{}
}

type repository struct {
	Books    []Book
	IDsCount int
}

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (r *repository) CreateBook(b Book) {
	r.IDsCount++
	b.ID = r.IDsCount
	r.Books = append(r.Books, b)
}

func (r *repository) UpdateBookByID(b Book) {
	for i, v := range r.Books {
		if v.ID == b.ID {
			r.Books[i] = b
			break
		}
	}

	return
}

func (r *repository) GetAllBooks() []Book {
	return r.Books
}

func (r *repository) GetByID(b Book) Book {
	for _, v := range r.Books {
		if v.ID == b.ID {
			b = v
			break
		}
	}

	return b
}

func (r *repository) DeleteByID(b Book) {
	for i, v := range r.Books {
		if v.ID == b.ID {
			r.Books = append(r.Books[:i], r.Books[i+1:]...)
			break
		}
	}
}

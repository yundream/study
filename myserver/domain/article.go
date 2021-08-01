package domain

type Article struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ArticleRepository interface {
	GetByID(id int64) (*Article, error)
	Fetch(offset, limit int) ([]*Article, error)
	Create(article *Article) (*Article, error)
	Update(id int64, article *Article) (*Article, error)
	Delete(id int64) error
}

type ArticleUseCase interface {
	GetByID(id int64) (res *Article, err error)
	Fetch(offset, limit int) (res []*Article, err error)
	Create(article *Article) (*Article, error)
	Update(id int64, article *Article) (*Article, error)
	Delete(id int64) error
}

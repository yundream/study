package mysql

import (
	"fmt"

	model "joinc.co.kr/study/myserver/domain"
)

func NewMySQLRepository() (*MySQLRepository, error) {
	return &MySQLRepository{DBEngine: "mysql"}, nil
}

type MySQLRepository struct {
	DBEngine string
}

func (m *MySQLRepository) GetByID(id int64) (*model.Article, error) {
	fmt.Println(">> ", m.DBEngine)
	return &model.Article{
		Title:    "MySQL 사용법",
		Body:     "MySQL은 쉽게 사용 할 수 있습니다.\n정말로",
		UserName: "yundream",
	}, nil
}

func (m *MySQLRepository) Fetch(offset, limit int) ([]*model.Article, error) {
	fmt.Println(">> ", m.DBEngine)
	return []*model.Article{
		{
			Title:    "MySQL 사용법",
			Body:     "MySQL은 쉽게 사용 할 수 있습니다.\n정말로",
			UserName: "yundream",
		},
		{
			Title:    "GoLang의 미래",
			Body:     "GoLang의 미래는 밝아보입니다....",
			UserName: "yundream",
		},
	}, nil
}

func (m *MySQLRepository) Create(article *model.Article) (*model.Article, error) {
	return article, nil
}

func (m *MySQLRepository) Update(id int64, article *model.Article) (*model.Article, error) {
	return &model.Article{Title: "Update"}, nil
}

func (m *MySQLRepository) Delete(id int64) error {
	fmt.Println("삭제")
	return nil
}

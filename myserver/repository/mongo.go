package repository

/*
import (
	"fmt"

	"joinc.co.kr/study/sqlmock/model"
)

func NewMongoRepository() (*MongoRepository, error) {
	return &MongoRepository{DBEngine: "mongodb"}, nil
}

type MongoRepository struct {
	DBEngine string
}

func (m *MongoRepository) GetByID(id int64) (*model.Article, error) {
	fmt.Println(">> ", m.DBEngine)
	return &model.Article{
		Title:    "MySQL 사용법",
		Body:     "MySQL은 쉽게 사용 할 수 있습니다.\n정말로",
		UserName: "yundream",
	}, nil
}

func (m *MongoRepository) Fetch(offset, limit int) ([]*model.Article, error) {
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

func (m *MongoRepository) Create(article *model.Article) (*model.Article, error) {
	return article, nil
}

func (m *MongoRepository) Update(id int64, article *model.Article) (*model.Article, error) {
	return &model.Article{Title: "Update"}, nil
}

func (m *MongoRepository) Delete(id int64) error {
	fmt.Println("삭제")
	return nil
}
*/

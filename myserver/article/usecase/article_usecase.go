package usecase

import "joinc.co.kr/study/myserver/domain"

type articleUsecase struct {
	articleRepo domain.ArticleRepository
}

func NewArticleUseCase(a domain.ArticleRepository) domain.ArticleUseCase {
	return &articleUsecase{
		articleRepo: a,
	}
}

func (a *articleUsecase) GetByID(id int64) (*domain.Article, error) {
	res, err := a.articleRepo.GetByID(id)
	return res, err
}
func (a *articleUsecase) Fetch(offset, limit int) ([]*domain.Article, error) {
	return a.articleRepo.Fetch(offset, limit)
}

func (a *articleUsecase) Create(article *domain.Article) (*domain.Article, error) {
	return a.articleRepo.Create(article)
}

func (a *articleUsecase) Update(id int64, article *domain.Article) (*domain.Article, error) {
	return a.articleRepo.Update(id, article)
}

func (a *articleUsecase) Delete(id int64) error {
	return a.articleRepo.Delete(id)
}

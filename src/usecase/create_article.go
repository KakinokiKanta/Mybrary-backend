package usecase

import articleDomain "github.com/KakinokiKanta/Mybrary-backend/domain"

type CreateArticleUseCase struct {
	articleRepo articleDomain.ArticleRepository
}

func NewCreateArticleUseCase (articleRepo articleDomain.ArticleRepository) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepo: articleRepo,
	}
}

type CreateArticleInputDto struct {
	UserID      string
	Url         string
	Title       string
	Description string
	Tags        []articleDomain.ArticleTag
}

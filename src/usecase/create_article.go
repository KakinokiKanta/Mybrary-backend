package usecase

import articleDomain "github.com/KakinokiKanta/Mybrary-backend/domain"

type CreateArticleUseCase struct {
	articleRepo articleDomain.ArticleRepository
}

type CreateArticleInputDto struct {
	UserID      string `json:"user_id" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Tags        []articleDomain.ArticleTag `json: tags`
}

type ArticleTag struct {
	TagName string `json:"tag_name" binding:"required"`
}

func NewCreateArticleUseCase (articleRepo articleDomain.ArticleRepository) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepo: articleRepo,
	}
}

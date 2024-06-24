package article

import articleDomain "github.com/KakinokiKanta/Mybrary-backend/domain"

type SaveArticleUseCase struct {
	articleRepo articleDomain.ArticleRepository
}

func NewSaveArticleUseCase (articleRepo articleDomain.ArticleRepository) *SaveArticleUseCase {
	return &SaveArticleUseCase{
		articleRepo: articleRepo,
	}
}

type SaveArticleInputDto struct {
	UserID      string
	Url         string
	Title       string
	Description string
	Tags        []articleDomain.ArticleTag
}

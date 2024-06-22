package article

import articleDomain "github.com/KakinokiKanta/Mybrary-backend/domain/article"

type SaveArticleUseCase struct {
	articleRepo articleDomain.ArticleRepository
}

func NewSaveArticleUseCase (articleRepo articleDomain.ArticleRepository) *SaveArticleUseCase {
	return &SaveArticleUseCase{
		articleRepo: articleRepo,
	}
}

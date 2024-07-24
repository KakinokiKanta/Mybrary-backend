package usecase

import "github.com/KakinokiKanta/Mybrary-backend/domain"

type CreateArticleUseCase struct {
	articleRepo domain.ArticleRepository
}

type CreateArticleInputDto struct {
	UserID      string `json:"user_id" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Tags        []string `json:"tags"`
}

type CreateArticleOutputDto struct {
	ArticleID   domain.ArticleID `json:"id"`
	UserID      string `json:"user_id"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        []ArticleTag `json:"tags"`
}

type ArticleTag struct {
	TagID string `json:"tag_id"`
	TagName string `json:"tag_name"`
	UsedNum int `json:"used_num"`
}

func NewCreateArticleUseCase (articleRepo domain.ArticleRepository) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepo: articleRepo,
	}
}

func (uc CreateArticleUseCase) Execute(input CreateArticleInputDto) (*CreateArticleOutputDto, error) {
	return &CreateArticleOutputDto{
		ArticleID: domain.ArticleID("aa"),
		UserID: "bb",
		Url: "cc",
		Title: "dd",
		Description: "ee",
		Tags: []ArticleTag{
			{
				TagID: "ff",
				TagName: "gg",
				UsedNum: 11,
			},
		},
	}, nil
}

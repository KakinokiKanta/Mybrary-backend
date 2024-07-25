package usecase

import (
	"database/sql"
	"errors"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type CreateArticleUseCase struct {
	articleRepo domain.ArticleRepository
	tagRepo domain.ArticleTagRepository
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
	// ArticleTagドメインを生成
	var articleTagList = []domain.ArticleTag{}
	var outputTagList = []ArticleTag{}
	for _, tagName := range input.Tags {
		articleTag, err := domain.NewArticleTag(domain.UserID(input.UserID), tagName)
		if err != nil {
			return nil, err
		}

		// 同一名のタグがDB内に存在しないかチェック
		var savedTag domain.ArticleTag
		_, err = uc.tagRepo.FindByName(articleTag.TagName())
		if err == nil {
			// ArticleTagの情報更新
			savedTag, err = uc.tagRepo.UpdateNum(tagName)
			if err != nil {
				return nil, err
			}
		} else if err != sql.ErrNoRows {
			// 単純なエラー時
			return nil, err
		} else {
			// ArticleTagの永続化
			savedTag, err = uc.tagRepo.Create(*articleTag)
			if err != nil {
				return nil, err
			}
		}

		articleTagList = append(articleTagList, savedTag)
		outputTagList = append(outputTagList, ArticleTag{
			TagID: string(savedTag.ID()),
			TagName: savedTag.TagName(),
			UsedNum: savedTag.UsedNum(),
		})
	}

	// Articleドメインを生成
	article, err := domain.NewArticle(domain.UserID(input.UserID), input.Url, input.Title, input.Description, articleTagList)
	if err != nil {
		return nil, err
	}

	// 同一URLを持つ記事がDB内に存在しないかチェック
	_, err = uc.articleRepo.FindByUrl(article.URL())
	if err == nil {
		return nil, errors.New("this URL is already registered")
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	// Articleの永続化
	createdArticle, err := uc.articleRepo.Create(*article)
	if err != nil {
		return nil, err
	}

	return &CreateArticleOutputDto{
		ArticleID: createdArticle.ID(),
		UserID: string(createdArticle.UserID()),
		Url: createdArticle.URL(),
		Title: createdArticle.Title(),
		Description: createdArticle.Description(),
		Tags: outputTagList,
	}, nil
}

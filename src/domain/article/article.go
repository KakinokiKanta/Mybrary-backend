package article

import "github.com/oklog/ulid/v2"

type Article struct {
	id          string
	userID      string
	url         string
	title       string
	description string
	tags        []string
}

func NewArticle(
	userID string,
	url string,
	title string,
	description string,
	tags []string,
) (*Article, error) {
	// ulidパッケージでULIDを生成し、string型に変換
	id := ulid.Make().String()

	// TODO: 各種項目のバリデーション

	return &Article{
		id: id,
		userID: userID,
		url: url,
		title: title,
		description: description,
		tags: tags,
	}, nil
}

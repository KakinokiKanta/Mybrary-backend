package article

type ArticleRepository interface {
	Create(article Article) error
	FindById(id string) (Article, error)
	FindAll() ([]Article, error)
	Update(article Article) error
	Delete(article Article) error
}
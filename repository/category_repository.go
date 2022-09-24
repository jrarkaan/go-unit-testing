package repository

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

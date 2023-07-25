package repository

import "golang-unit-test-basic/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

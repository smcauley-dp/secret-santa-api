package repositories

import "github.com/KoseSoftware/secret-santa-api/models"

type ListerRepository interface {
	Create(l models.List) (id string, err error)
	FindAll(email string) (items []models.List, err error)
	FindByID(id string) (item models.List, err error)
	DeleteByID(id string) (rowsAffected int64, err error)
}

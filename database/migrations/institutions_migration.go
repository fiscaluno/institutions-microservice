package migrations

import (
	model "github.com/fiscaluno/institutions-microservice/models"
)

var Institution = model.Institution{}

func init() {
	Create(&Institution)
}

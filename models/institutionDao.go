package models

import(
    "github.com/fiscaluno/institutions-microservice/database"
    _ "github.com/jinzhu/gorm"
)

var db = database.GetInstance()

package db

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

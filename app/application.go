package app

import (
	"github.com/gandra/bookstore/usersapigorm/datasources/postgresql/users_db"
	"github.com/gandra/bookstore/usersapigorm/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	users_db.InitDatabase()
	mapUrls()

	logger.Info("about to start application...")
	router.Run(":8080")
}

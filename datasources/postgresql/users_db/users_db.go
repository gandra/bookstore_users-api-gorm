package users_db

import (
	"fmt"
	"github.com/gandra/bookstore/usersapigorm/datasources/postgresql/db"
	"github.com/gandra/bookstore/usersapigorm/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	db_users_username = "db_users_username"
	db_users_password = "db_users_password"
	db_users_host     = "db_users_host"
	db_users_schema   = "db_users_schema"
)

var (
	username = os.Getenv(db_users_username)
	password = os.Getenv(db_users_password)
	host     = os.Getenv(db_users_host)
	schema   = os.Getenv(db_users_schema)
)

func InitDatabase() {
	// TOD: Fix temp hack to avoid playing with env vars
	username = "postgres"
	password = "mysecret"
	host = "localhost"
	schema = "bookings_users"

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", host, username, password, schema)
	db.Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected!")
	db.Client.AutoMigrate(&users.User{})
	fmt.Println("Migrated DB")

}

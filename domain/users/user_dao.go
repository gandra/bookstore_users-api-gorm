package users

import (
	"fmt"
	"github.com/gandra/bookstore/usersapigorm/datasources/postgresql/db"
	"github.com/gandra/bookstore/usersapigorm/logger"
	"github.com/gandra/bookstore/usersapigorm/utils/errors"
)

const (
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users where status = ?;"
)

func (user *User) Get() *errors.RestErr {
	err := db.Client.Find(user, user.Id).Error
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("Not found user for ids %s", user.Id))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	err := db.Client.Create(&user).Error
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	err := db.Client.Save(user).Error
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	err := db.Client.Delete(user).Error
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	results := make([]User, 0)
	rows, err := db.Client.Raw(queryFindUserByStatus, status).Rows()
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	return results, nil
}

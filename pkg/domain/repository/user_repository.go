package repository

import "go-labs-sql/pkg/domain"

type UserRepository interface {

	Create(data *domain.UserData) (domain.User,error)

}

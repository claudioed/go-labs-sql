package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"go-labs-sql/pkg/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func (repo PostgresUserRepository) Create(data *domain.UserData) (domain.User, error) {
	userId, err := uuid.NewRandom()
	if err != nil {
		logger.Error("Error to generate UUID")
	}
	user := domain.User{
		Name:  data.Name,
		Email: data.Email,
		Id:    userId.String(),
	}
	_, dbErr := repo.db.Exec("INSERT INTO users_data ( id,name,email) VALUES($1, $2, $3)", user.Id, user.Name, user.Email)
	if dbErr != nil {
		logger.Error("Error to persist user",dbErr)
		return user, errors.New("ERROR_IN_DB")
	}
	logger.Info("User created successfully!!!")
	return user, nil
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	logger.Info("Creating UserRepository Instance (PostgreSQL)")
	return &PostgresUserRepository{db}
}

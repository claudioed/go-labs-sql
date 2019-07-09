package handlers

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"go-labs-sql/pkg/domain"
	"go-labs-sql/pkg/domain/repository"
	"net/http"
)

type UserHandler struct {
	Repository *repository.PostgresUserRepository
}

func (handler *UserHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data *domain.UserData
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			w.WriteHeader(500)
			return
		}
		logger.WithFields(logger.Fields{
			"name": data.Name,
			"email" : data.Email,
		}).Info("Creating New User")

		defer r.Body.Close()

		user,dataErr := handler.Repository.Create(data)

		if dataErr != nil {
			w.WriteHeader(500)
			str := `{"error": "error to create user"}`
			errorMessage, err :=json.Marshal([]byte(str))
			if err != nil{
				logger.Error("Error to persist user")
			}
			_, _ = w.Write(errorMessage)
			return
		}
		userJson, errJson := json.Marshal(user)
		if errJson != nil {
			w.WriteHeader(500)
			str := `{"error": "error to format json"}`
			errorMessage, err :=json.Marshal([]byte(str))
			if err != nil{
				logger.Error("error to format json")
			}
			_, _ = w.Write(errorMessage)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(userJson)

	})
}

func NewUserHandler(userRepository *repository.PostgresUserRepository) *UserHandler {
	logger.Info("Creating User Handler Instance")
	return &UserHandler{userRepository}
}
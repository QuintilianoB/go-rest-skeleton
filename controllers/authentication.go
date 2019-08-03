package controllers

import (
	"go-rest-skeleton/models"
	"go-rest-skeleton/repository"
	"go-rest-skeleton/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func (c Controller) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		user, err := utils.ValidateUser(r)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, err)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			return
		}

		// Convertendo de bytes para hash.
		user.Password = string(hash)

		userRepo := repository.UserRepository{}
		err = userRepo.CreateUser(&user)

		if err != nil {
			if err == models.UserAlreadyExist {
				utils.SendError(w, http.StatusNotModified, err)
				return
			} else {
				utils.SendError(w, http.StatusInternalServerError, err)
				return
			}
		}
		log.Printf("User %s created.\n", user.User)
		w.WriteHeader(http.StatusCreated)
		utils.SendSuccess(w, nil)
	}
}

func (c Controller) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var jwt models.JWT

		user, err := utils.ValidateUser(r)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, err)
			return
		}

		password := user.Password

		userRepo := repository.UserRepository{}
		err = userRepo.FindUser(&user)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, models.InvalidCredentials)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.InvalidCredentials)
			return
		}

		jwt.Token = token

		log.Printf("User %s authenticated.\n", user.User)
		utils.SendSuccess(w, jwt)
	}
}

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"iupDB/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnv(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendError(writer http.ResponseWriter, status int, error error) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-type", "application/json")

	if error == nil {
		json.NewEncoder(writer)
	} else {
		err := json.NewEncoder(writer).Encode(error)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendSuccess(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-type", "application/json")

	if data == nil {
		json.NewEncoder(writer)
	} else {
		err := json.NewEncoder(writer).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func ValidateUser(r *http.Request) (models.User, error) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return user, models.InvalidJson
	}

	if user.User == "" || user.Password == "" {
		return user, models.InvalidCredentials
	}

	return user, nil
}

func GenerateToken(user models.User) (string, error) {
	secret := GetEnv("secret", "")

	if secret == "" {
		log.Fatal("Secret not configured. Check .env file.")
	}

	tokenTtl := GetEnv("tokenTtl", "900")
	ttl, err := strconv.Atoi(tokenTtl)

	if err != nil {
		log.Fatal("Invalid tokenTtl. Check .env file.")
	}

	exp := time.Now().Add(time.Duration(ttl) * time.Second)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.User,
		"iss":  "iupDB",
		"exp":  exp,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			SendError(w, http.StatusUnauthorized, models.InvalidToken)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		secret := os.Getenv("secret")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, models.InvalidToken
				}

				return []byte(secret), nil
			})

			if err != nil {
				SendError(w, http.StatusUnauthorized, err)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				exp := fmt.Sprintf("%s", claims["exp"])

				// Não rela.
				layout := "2006-01-02T15:04:05.000000000-07:00"

				t, err := time.Parse(layout, exp)
				if err != nil {
					SendError(w, http.StatusInternalServerError, err)
					return
				}

				if time.Now().After(t) {
					SendError(w, http.StatusUnauthorized, models.TokenExpired)
					return
				} else {
					next.ServeHTTP(w, r)
				}
			} else {
				SendError(w, http.StatusUnauthorized, err)
				return
			}
		} else {
			SendError(w, http.StatusBadRequest, models.InvalidToken)
			return
		}
	})
}

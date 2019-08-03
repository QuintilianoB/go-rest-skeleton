package controllers

import (
	"go-rest-skeleton/utils"
	"net/http"
)

func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SendSuccess(w, "ola")
	}
}

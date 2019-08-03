package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-skeleton/utils"
	"log"
	"net/http"
)

type Controller struct{}

func Server() {

	controller := Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.SignUp()).Methods("POST")
	router.HandleFunc("/login", controller.Login()).Methods("POST")
	router.HandleFunc("/protected", utils.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")

	host := utils.GetEnv("serverAddr", "127.0.0.1")
	port := utils.GetEnv("serverPort", "8000")
	serverUrl := fmt.Sprint(host + ":" + port)

	fmt.Println("Start serving on", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, router))
}

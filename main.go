package main

import (
	"github.com/subosito/gotenv"
	"go-rest-skeleton/controllers"
	"log"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	controllers.Server()

}

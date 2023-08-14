package main

import (
	"github.com/CanadianCommander/devwiki/services/ui/internal/web"
	"log"
)

func main() {

	router := web.SetupPages()

	err := router.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal("Unexpected error in Gin webserver! ", err)
	}
}

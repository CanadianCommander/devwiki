package main

import (
	"github.com/CanadianCommander/devwiki/services/ui/internal/pages"
	"log"
)

func main() {

	router := pages.SetupPages()

	err := router.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal("Unexpected error in Gin webserver! ", err)
	}
}

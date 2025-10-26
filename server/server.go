package main

import (
	"log"
	"os"
	"workbench/reiverrbridge/routes"
)

const defaultPort = "6060"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := routes.InitRouter()

	log.Printf("Server run on http://localhost:%s/", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("can't start server:", err)
	}
}

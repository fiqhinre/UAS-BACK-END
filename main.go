package main

import (
	"log"
	"unairsatu_v2/config"
	"unairsatu_v2/routes"
)

func main() {
	// Hubungkan ke MongoDB
	config.ConnectDB()

	// Inisialisasi routing
	r := routes.SetupRoutes()

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

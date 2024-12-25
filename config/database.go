package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Konfigurasi URI MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://fiqhi:Mieayam2porsi!@pertama.a8bqn5s.mongodb.net/")

	// Membuat koneksi ke MongoDB
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi konteks dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Koneksi ke server MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Pilih database
	DB = client.Database("unairsatu_v2")
	log.Println("Koneksi ke MongoDB berhasil!")
}

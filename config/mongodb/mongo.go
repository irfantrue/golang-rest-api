package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	mongoURI   = "mongodb://localhost:27017/development" // Ganti dengan URI MongoDB Anda
	timeoutSec = 10
)

// Connect digunakan untuk membuat koneksi ke MongoDB.
func Connect() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	newClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Tes koneksi
	err = newClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	client = newClient
	fmt.Println("Koneksi MongoDB berhasil dibuat")
	return client, nil
}

// Disconnect digunakan untuk memutus koneksi dari MongoDB.
func Disconnect() {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			fmt.Println("Gagal memutus koneksi MongoDB:", err)
		} else {
			fmt.Println("Koneksi MongoDB berhasil diputus")
		}
		client = nil
	}
}

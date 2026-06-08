package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		fmt.Println("Error creating Firestore client:", err)
		return
	}

	fmt.Println("Firestore client created successfully")

	defer client.Close()
}

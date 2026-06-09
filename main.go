package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/danieldesira/firestore-rooted-bulk-inserter/lib"
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

	file, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Unable to open questions.csv:", err)
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		fmt.Println(lib.MapEntryToQuestion(record))
	}

	fmt.Println("Firestore client created successfully")

	defer file.Close()
	defer client.Close()
}

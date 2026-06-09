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

	insertedCount := 0
	for _, record := range records {
		question := lib.MapEntryToQuestion(record)

		_, err := client.Collection("question").NewDoc().Set(ctx, question)
		if err != nil {
			fmt.Printf("Error inserting question %v: %v\n", question, err)
			continue
		}
		insertedCount++
	}

	fmt.Printf("Successfully inserted %d questions into Firestore\n", insertedCount)

	defer file.Close()
	defer client.Close()
}

package lib

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

func ImportCsv() {
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
		question := MapEntryToQuestion(record)

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

func ExportCsv() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		fmt.Println("Error creating Firestore client:", err)
		return
	}
	defer client.Close()

	// Fetch all questions from Firestore
	docs, err := client.Collection("question").Documents(ctx).GetAll()
	if err != nil {
		fmt.Println("Error fetching questions from Firestore:", err)
		return
	}

	// Create or overwrite the CSV file
	file, err := os.Create("questions.csv")
	if err != nil {
		fmt.Println("Error creating questions.csv:", err)
		return
	}
	defer file.Close()

	// Write CSV records
	writer := csv.NewWriter(file)
	defer writer.Flush()

	exportedCount := 0
	for _, doc := range docs {
		var question Question
		if err := doc.DataTo(&question); err != nil {
			fmt.Printf("Error mapping document %s: %v\n", doc.Ref.ID, err)
			continue
		}

		entry := MapQuestionToEntry(question)
		if err := writer.Write(entry); err != nil {
			fmt.Printf("Error writing question %v: %v\n", question, err)
			continue
		}
		exportedCount++
	}

	fmt.Printf("Successfully exported %d questions to questions.csv\n", exportedCount)
}

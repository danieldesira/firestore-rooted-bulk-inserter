package lib

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

func initFirestore() (context.Context, *firestore.Client, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}

func openFile() (*os.File, error) {
	file, err := os.Open("questions.csv")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ImportCsv() {
	ctx, client, err := initFirestore()
	if err != nil {
		fmt.Println("Firestore error: ", err)
		return
	}

	file, err := openFile()
	if err != nil {
		fmt.Println("File error: ", err)
		return
	}

	defer file.Close()
	defer client.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	docs, err := client.Collection("question").Documents(ctx).GetAll()
	if err != nil {
		fmt.Println("Error fetching questions from Firestore:", err)
		return
	}

	for _, doc := range docs {
		client.Collection("question").Doc(doc.Ref.ID).Delete(ctx)
	}

	insertedCount := 0
	for _, record := range records[1:] {
		question := MapEntryToQuestion(record)

		_, err := client.Collection("question").NewDoc().Set(ctx, question)
		if err != nil {
			fmt.Printf("Error inserting question %v: %v\n", question, err)
			continue
		}
		insertedCount++
	}

	fmt.Printf("Successfully inserted %d questions into Firestore\n", insertedCount)

}

func createFile() (*os.File, error) {
	file, err := os.Create("questions.csv")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ExportCsv() {
	ctx, client, err := initFirestore()
	if err != nil {
		fmt.Println("Firestore error: ", err)
		return
	}

	file, err := createFile()
	if err != nil {
		fmt.Println("File error: ", err)
		return
	}

	defer file.Close()
	defer client.Close()

	docs, err := client.Collection("question").Documents(ctx).GetAll()
	if err != nil {
		fmt.Println("Error fetching questions from Firestore:", err)
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	exportedCount := 0
	writer.Write([]string{"Subject", "Tag", "Question", "Visual", "Option 1 (Correct)", "Option 2", "Option 3", "Option 4"})
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/danieldesira/firestore-rooted-bulk-inserter/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	fmt.Println("Rooted Bulk Importer")
	fmt.Println("--------------------")
	fmt.Println()

	fmt.Println("Enter 1: Export")
	fmt.Println("Enter 2: Import")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch strings.TrimSpace(input) {
	case "1":
		lib.ExportCsv()
	case "2":
		lib.ImportCsv()
	default:
		fmt.Println("No option selected...")
	}
}

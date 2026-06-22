# Firestore Rooted Bulk Inserter

A Go application to bulk insert questions into Firestore from a CSV file.

## Prerequisites

### Google Cloud CLI

Install the Google Cloud CLI to authenticate with your Firestore database. See the [official installation guide](https://cloud.google.com/sdk/docs/install).

After installation, authenticate with:
```bash
gcloud auth login
```

### Environment Variables

Create a `.env` file in the project root with the following:
```
FIRESTORE_PROJECT_ID=your-project-id
```

## Running the Program

1. Ensure you have a `questions.csv` file in the project root.
2. Run the program:
```bash
go run main.go
```

## CSV Structure

The `questions.csv` file should have the following structure:

```
SubjectId,Tag,Question,Visual,Option1,Option2,Option3,Option4
```

### Columns:

- **subjectId**: Identifier for the subject/category
- **tag**: Tag or topic label for the question
- **question**: The question text
- **visual**: Base64 image
- **option1**: First answer option (marked as correct)
- **option2**: Second answer option
- **option3**: Third answer option
- **option4**: Fourth answer option

### Example:

```csv
math,algebra,What is 2+2?,4,3,5,6
science,physics,What is the speed of light?,299792458,300000000,100000000,50000000
```

The first option (option1) is automatically marked as the correct answer.

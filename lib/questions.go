package lib

type Question struct {
	SubjectId string `firestore:"subjectId"`
	Tag       string `firestore:"tag"`
	Question  string `firestore:"question"`
	Options   [4]struct {
		Answer    string `firestore:"answer"`
		IsCorrect bool   `firestore:"isCorrect"`
	} `firestore:"options"`
}

func MapEntryToQuestion(entry []string) Question {
	var question Question
	question.SubjectId = entry[0]
	question.Tag = entry[1]
	question.Question = entry[2]
	for i := 3; i < len(entry); i++ {
		question.Options[i-3].Answer = entry[i]
	}
	question.Options[0].IsCorrect = true
	return question
}

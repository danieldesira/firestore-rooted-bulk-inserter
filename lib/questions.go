package lib

type Question struct {
	SubjectId string `firestore:"subjectId"`
	Tag       string `firestore:"tag"`
	Question  string `firestore:"question"`
	Visual    string `firestore:"visual"`
	Options   [4]struct {
		Answer    string `firestore:"answer"`
		IsCorrect bool   `firestore:"isCorrect"`
	} `firestore:"options"`
}

const subjectIndex = 0
const tagIndex = 1
const questionIndex = 2
const visualIndex = 3
const optionIndex = 4

func MapEntryToQuestion(entry []string) Question {
	var question Question
	question.SubjectId = entry[subjectIndex]
	question.Tag = entry[tagIndex]
	question.Question = entry[questionIndex]
	question.Visual = entry[visualIndex]
	for i := optionIndex; i < len(entry); i++ {
		question.Options[i-optionIndex].Answer = entry[i]
	}
	question.Options[0].IsCorrect = true
	return question
}

func MapQuestionToEntry(question Question) []string {
	entry := [8]string{
		question.SubjectId,
		question.Tag,
		question.Question,
		question.Visual,
		question.Options[0].Answer,
		question.Options[1].Answer,
		question.Options[2].Answer,
		question.Options[3].Answer,
	}
	return entry[:]
}

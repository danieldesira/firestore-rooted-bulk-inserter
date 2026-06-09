package lib

type Question struct {
	subjectId string
	tag       string
	question  string
	options   [4]struct {
		answer    string
		isCorrect bool
	}
}

func MapEntryToQuestion(entry []string) Question {
	var question Question
	question.subjectId = entry[0]
	question.tag = entry[1]
	question.question = entry[2]
	for i := 3; i < len(entry); i++ {
		question.options[i-3].answer = entry[i]
	}
	question.options[0].isCorrect = true
	return question
}

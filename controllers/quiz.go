package controllers

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/eas-kaine/discord-bot/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func HandleQuiz(s *discord.Session, m *discord.MessageCreate) {
	cmd := strings.Fields(m.Content)
	// !quiz
	switch cmd[1] {
	case "create":
		createQuiz(s)
	case "add_question":
		addQuestion(s)
	case "start":
		startQuiz(s)
	case "end":
		endQuiz(s)
	case "delete_question":
		deleteQuestion(s)
	case "answer":
		answerQuestion(s)
	}
}

func createQuiz(s *discord.Session) {
	quiz := models.Quiz{
		Name: "Test Quiz",
		Subject: "Testing",
		Question: models.Question{
			Q: "Is this a test?",
			Answer: "yes",
		}}
	// result := DB.Create(&quiz)

	fmt.Print(DB.First(&quiz))

	//fmt.Println(result.RowsAffected)

	fmt.Println(quiz.Name)

	// quiz := models.Quiz{
	// 	Name: "Test Quiz",
	// 	Subject: "Testing",
	// 	Questions: models.Questions{models.Question{
	// 		Q: "Is this a test?",
	// 		Answer: "yes",
	// 	}}}
	// DB.Create(&quiz)
}

func addQuestion(s *discord.Session) {

}

func startQuiz(s *discord.Session) {

}

func endQuiz(s *discord.Session) {

}

func deleteQuestion(s *discord.Session) {

}

func answerQuestion(s *discord.Session) {

}

// func listQuizzes(s *discord.Session) {
	
// }
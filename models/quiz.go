package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Name string
	Subject string
	QuestionID int
	Question
	Count int
}

type Question struct {
	ID int
	Q string
	Answer string
}

// type Quiz struct {
// 	gorm.Model
// 	Name string
// 	Subject string
// 	QuestionID int
// 	Questions
// 	Count int
// }

// type Question struct {
// 	ID int
// 	Q string
// 	Answer string
// }

// type Questions []Question
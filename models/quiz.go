package models

import "gorm.io/gorm"

type Quiz struct {
    gorm.Model
    Title     string     `json:"title" binding:"required"`
    LessonID  uint       `json:"lesson_id"`
    Questions []Question `json:"questions"`
}

type Question struct {
    gorm.Model
    QuizID        uint     `json:"quiz_id"`
    Text          string   `json:"text" binding:"required"`
    Options       string   `json:"options"`
    CorrectAnswer string   `json:"correct_answer" binding:"required"`
}

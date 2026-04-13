package models

import "gorm.io/gorm"

type QuizResult struct {
    gorm.Model
    UserID    uint `json:"user_id"`
    QuizID    uint `json:"quiz_id"`
    Score     int  `json:"score"`
    Correct   int  `json:"correct"`
    Total     int  `json:"total"`
}

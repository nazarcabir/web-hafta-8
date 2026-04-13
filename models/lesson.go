package models

import "gorm.io/gorm"

type Lesson struct {
    gorm.Model
    Title     string `json:"title" binding:"required"`
    Content   string `json:"content"`
    CourseID  uint   `json:"course_id"`
    Order     int    `json:"order"`
}

package models

import "gorm.io/gorm"

type Course struct {
    gorm.Model
    Title       string   `json:"title" binding:"required"`
    Description string   `json:"description"`
    Category    string   `json:"category"`
    TeacherID   uint     `json:"teacher_id"`
    Lessons     []Lesson `json:"lessons,omitempty"`
}

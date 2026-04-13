package database

import (
    "golearn/models"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("golearn.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connection established and migrating models...")
    
    // Auto Migrate
    DB.AutoMigrate(
        &models.User{},
        &models.Course{},
        &models.Lesson{},
        &models.Quiz{},
        &models.Question{},
        &models.Progress{},
        &models.Enrollment{},
        &models.QuizResult{},
    )
}

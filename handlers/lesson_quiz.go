package handlers

import (
    "golearn/database"
    "golearn/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// @Summary Kursa Ait Dersleri Listele
// @Tags Lesson
// @Param id path int true "Kurs ID"
// @Success 200 {array} models.Lesson
// @Router /courses/{id}/lessons [get]
func GetLessons(c *gin.Context) {
    var lessons []models.Lesson
    database.DB.Where("course_id = ?", c.Param("id")).Order("\"order\" asc").Find(&lessons)
    c.JSON(http.StatusOK, lessons)
}

// @Summary Kursa Yeni Ders Ekle
// @Tags Lesson
// @Security BearerAuth
// @Param id path int true "Kurs ID"
// @Param body body models.Lesson true "Ders Bilgileri"
// @Success 201 {object} models.Lesson
// @Router /courses/{id}/lessons [post]
func CreateLesson(c *gin.Context) {
    var lesson models.Lesson
    if err := c.ShouldBindJSON(&lesson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    courseID, _ := strconv.Atoi(c.Param("id"))
    lesson.CourseID = uint(courseID)
    database.DB.Create(&lesson)
    c.JSON(http.StatusCreated, lesson)
}

// @Summary Derse Ait Quiz Getir
// @Tags Quiz
// @Param id path int true "Ders ID"
// @Success 200 {object} models.Quiz
// @Router /lessons/{id}/quiz [get]
func GetQuiz(c *gin.Context) {
    var quiz models.Quiz
    if err := database.DB.Preload("Questions").Where("lesson_id = ?", c.Param("id")).First(&quiz).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz bulunamadı"})
        return
    }
    c.JSON(http.StatusOK, quiz)
}

// @Summary Yeni Quiz Oluştur
// @Tags Quiz
// @Security BearerAuth
// @Param id path int true "Ders ID"
// @Param body body models.Quiz true "Quiz Bilgileri"
// @Success 201 {object} models.Quiz
// @Router /lessons/{id}/quiz [post]
func CreateQuiz(c *gin.Context) {
    var quiz models.Quiz
    if err := c.ShouldBindJSON(&quiz); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    lessonID, _ := strconv.Atoi(c.Param("id"))
    quiz.LessonID = uint(lessonID)
    database.DB.Create(&quiz)
    c.JSON(http.StatusCreated, quiz)
}

type QuizSubmission struct {
    QuestionID uint   `json:"question_id"`
    Answer     string `json:"answer"`
}

// @Summary Quiz Çöz ve Puanla
// @Tags Quiz
// @Security BearerAuth
// @Param id path int true "Quiz ID"
// @Param body body []QuizSubmission true "Cevaplar"
// @Success 200 {object} models.QuizResult
// @Router /quiz/{id}/submit [post]
func SubmitQuiz(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    quizID, _ := strconv.Atoi(c.Param("id"))

    var answers []QuizSubmission
    if err := c.ShouldBindJSON(&answers); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var questions []models.Question
    database.DB.Where("quiz_id = ?", quizID).Find(&questions)
    
    correct := 0
    for _, q := range questions {
        for _, a := range answers {
            if q.ID == a.QuestionID && q.CorrectAnswer == a.Answer {
                correct++
            }
        }
    }

    score := 0
    if len(questions) > 0 {
        score = (correct * 100) / len(questions)
    }

    result := models.QuizResult{
        UserID:  userID,
        QuizID:  uint(quizID),
        Score:   score,
        Correct: correct,
        Total:   len(questions),
    }
    database.DB.Create(&result)

    c.JSON(http.StatusOK, result)
}

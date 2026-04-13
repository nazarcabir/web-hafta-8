package handlers

import (
    "golearn/database"
    "golearn/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// @Summary Dersi Tamamla
// @Tags Progress
// @Security BearerAuth
// @Param id path int true "Ders ID"
// @Success 200 {object} map[string]string
// @Router /lessons/{id}/complete [post]
func CompleteLesson(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    idStr := c.Param("id")
    lessonID, _ := strconv.Atoi(idStr)
    
    var lesson models.Lesson
    if err := database.DB.First(&lesson, lessonID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ders bulunamadı"})
        return
    }

    progress := models.Progress{
        UserID:    userID,
        LessonID:  uint(lessonID),
        CourseID:  lesson.CourseID,
        Completed: true,
    }
    database.DB.Where("user_id = ? AND lesson_id = ?", userID, lessonID).FirstOrCreate(&progress)
    c.JSON(http.StatusOK, gin.H{"message": "Ders tamamlandı"})
}

// @Summary İlerlememi Getir
// @Tags Progress
// @Security BearerAuth
// @Success 200 {array} map[string]interface{}
// @Router /my/progress [get]
func GetProgress(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    var results []struct {
        CourseID   uint    `json:"course_id"`
        Title      string  `json:"title"`
        Percentage float64 `json:"percentage"`
    }
    database.DB.Table("progresses").
        Select("progresses.course_id, courses.title, (CAST(COUNT(progresses.id) AS FLOAT) / (SELECT COUNT(*) FROM lessons WHERE lessons.course_id = progresses.course_id)) * 100 as percentage").
        Joins("JOIN courses ON courses.id = progresses.course_id").
        Where("progresses.user_id = ? AND progresses.completed = ?", userID, true).
        Group("progresses.course_id, courses.title").
        Scan(&results)
    c.JSON(http.StatusOK, results)
}

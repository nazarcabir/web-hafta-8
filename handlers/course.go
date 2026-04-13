package handlers

import (
    "golearn/database"
    "golearn/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// @Summary Kursları Filtrele ve Listele
// @Tags Course
// @Param page query int false "Sayfa"
// @Param limit query int false "Limit"
// @Param category query string false "Kategori"
// @Param sort query string false "Sıralama (asc, desc)"
// @Success 200 {array} models.Course
// @Router /courses [get]
func GetCourses(c *gin.Context) {
    var courses []models.Course
    query := database.DB.Model(&models.Course{})

    if cat := c.Query("category"); cat != "" {
        query = query.Where("category = ?", cat)
    }

    sort := c.DefaultQuery("sort", "desc")
    query = query.Order("created_at " + sort)

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset := (page - 1) * limit

    query.Limit(limit).Offset(offset).Find(&courses)
    c.JSON(http.StatusOK, gin.H{"data": courses, "page": page, "limit": limit})
}

// @Summary Kurs Detayı
// @Tags Course
// @Success 200 {object} models.Course
// @Router /courses/{id} [get]
func GetCourse(c *gin.Context) {
    var course models.Course
    if err := database.DB.Preload("Lessons").First(&course, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kurs bulunamadı"})
        return
    }
    c.JSON(http.StatusOK, course)
}

// @Summary Yeni Kurs Oluştur
// @Tags Course
// @Security BearerAuth
// @Router /courses [post]
func CreateCourse(c *gin.Context) {
    var course models.Course
    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course.TeacherID = c.MustGet("user_id").(uint)
    database.DB.Create(&course)
    c.JSON(http.StatusCreated, course)
}

// @Summary Kurs Güncelle
// @Tags Course
// @Security BearerAuth
// @Router /courses/{id} [put]
func UpdateCourse(c *gin.Context) {
    var course models.Course
    if err := database.DB.First(&course, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kurs bulunamadı"})
        return
    }
    if course.TeacherID != c.MustGet("user_id").(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Yetkiniz yok"})
        return
    }
    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Save(&course)
    c.JSON(http.StatusOK, course)
}

// @Summary Kursu Sil
// @Tags Course
// @Security BearerAuth
// @Param id path int true "Kurs ID"
// @Success 200 {object} map[string]string
// @Router /courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
    var course models.Course
    if err := database.DB.First(&course, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kurs bulunamadı"})
        return
    }
    if course.TeacherID != c.MustGet("user_id").(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Yetkiniz yok"})
        return
    }
    database.DB.Delete(&course)
    c.JSON(http.StatusOK, gin.H{"message": "Kurs silindi"})
}

// @Summary Kursa Kaydol (Enrolment)
// @Tags Course
// @Security BearerAuth
// @Param id path int true "Kurs ID"
// @Success 201 {object} models.Enrollment
// @Router /courses/{id}/enroll [post]
func EnrollCourse(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)
    courseID, _ := strconv.Atoi(c.Param("id"))

    var course models.Course
    if err := database.DB.First(&course, courseID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kurs bulunamadı"})
        return
    }

    var enrollment models.Enrollment
    if err := database.DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrollment).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Zaten bu kursa kayıtlısınız"})
        return
    }

    enrollment = models.Enrollment{
        UserID:   userID,
        CourseID: uint(courseID),
    }
    database.DB.Create(&enrollment)
    c.JSON(http.StatusCreated, enrollment)
}

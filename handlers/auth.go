package handlers

import (
    "golearn/database"
    "golearn/middleware"
    "golearn/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

// @Summary Kullanıcı Kaydı
// @Tags Auth
// @Param body body interface{} true "Kayıt Bilgileri"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {
    var input struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
        Role     string `json:"role"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Role == "" { input.Role = "student" }

    hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hash),
        Role:     input.Role,
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "E-posta kullanımda"})
        return
    }

    token, _ := middleware.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusCreated, gin.H{"token": token, "user_id": user.ID, "role": user.Role})
}

// @Summary Kullanıcı Girişi
// @Tags Auth
// @Param body body interface{} true "Giriş Bilgileri"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz bilgiler"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz bilgiler"})
        return
    }

    token, _ := middleware.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID, "role": user.Role})
}

package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func TeacherOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "teacher" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Bu işlem için 'teacher' yetkisi gereklidir"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// @title GoLearn API
// @version 1.0
// @description Uzaktan Eğitim Platformu REST API
// @host localhost:8090
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
    "log"
    "golearn/database"
    "golearn/handlers"
    "golearn/middleware"
    "github.com/gin-gonic/gin"
    
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "golearn/docs"
)

func main() {
    // 1. Veritabanı Bağlantısı
    database.Connect()

    // 2. Gin Engine
    r := gin.Default()

    // 3. Global Middleware
    r.Use(middleware.CORSMiddleware())
    r.Use(middleware.RateLimitMiddleware())

    // 4. Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // 5. API Routes
    api := r.Group("/api")
    {
        // Auth (Public)
        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }

        // Protected Routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            // Courses
            courses := protected.Group("/courses")
            {
                courses.GET("", handlers.GetCourses)
                courses.GET("/:id", handlers.GetCourse)
                courses.POST("", middleware.TeacherOnly(), handlers.CreateCourse)
                courses.PUT("/:id", middleware.TeacherOnly(), handlers.UpdateCourse)
                courses.DELETE("/:id", middleware.TeacherOnly(), handlers.DeleteCourse)
                courses.POST("/:id/enroll", handlers.EnrollCourse)
                
                // Nested Lessons
                courses.GET("/:id/lessons", handlers.GetLessons)
                courses.POST("/:id/lessons", middleware.TeacherOnly(), handlers.CreateLesson)
            }

            // Quiz & Progress
            protected.GET("/lessons/:id/quiz", handlers.GetQuiz)
            protected.POST("/lessons/:id/quiz", handlers.CreateQuiz)
            protected.POST("/quiz/:id/submit", handlers.SubmitQuiz)
            protected.POST("/lessons/:id/complete", handlers.CompleteLesson)
            protected.GET("/my/progress", handlers.GetProgress)
        }

        // WebSocket Classroom
        api.GET("/ws/classroom/:courseId", middleware.AuthMiddleware(), handlers.ClassroomWS)
    }

    log.Println("Server starting on :8090...")
    if err := r.Run(":8090"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}

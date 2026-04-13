package middleware

import (
    "net/http"
    "sync"
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

var (
    limiters = make(map[string]*rate.Limiter)
    mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()
    l, exists := limiters[ip]
    if !exists {
        l = rate.NewLimiter(5, 10)
        limiters[ip] = l
    }
    return l
}

func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        limiter := getLimiter(ip)
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Çok fazla istek gönderildi"})
            c.Abort()
            return
        }
        c.Next()
    }
}

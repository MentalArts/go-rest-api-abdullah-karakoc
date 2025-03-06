package middlewares

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	limiter *rate.Limiter
	mu      sync.Mutex
}

func NewRateLimiter(rps int, burst int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		if !rl.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

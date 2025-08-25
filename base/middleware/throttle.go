package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	tokens     int
	lastRefill time.Time
	rate       int           // tokens per window
	window     time.Duration // time window
	maxTokens  int
	mu         sync.Mutex
}

func newRateLimiter(rate int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		tokens:     rate,
		lastRefill: time.Now(),
		rate:       rate,
		window:     window,
		maxTokens:  rate,
	}
}

type visitor struct {
	limiter  *rateLimiter
	lastSeen time.Time
}

func (rl *rateLimiter) allow() (bool, int, int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)

	// Refill tokens when the window passes
	if elapsed >= rl.window {
		rl.tokens = rl.maxTokens
		rl.lastRefill = now
	}

	remaining := rl.tokens
	if rl.tokens > 0 {
		rl.tokens--
		remaining = rl.tokens
		return true, remaining, 0
	}

	// Calculate retry-after
	retryAfter := int(rl.window.Seconds() - elapsed.Seconds())
	if retryAfter < 0 {
		retryAfter = 0
	}
	return false, 0, retryAfter
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

func getVisitor(ip string, limit int, window time.Duration) *rateLimiter {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("visitors", visitors)

	v, exists := visitors[ip]
	if !exists {
		rl := newRateLimiter(limit, window)
		visitors[ip] = &visitor{rl, time.Now()}
		return rl
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// Background cleanup for stale IPs
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// Throttle middleware (like Laravel throttle:60,1)
func Throttle(rule string) gin.HandlerFunc {
	parts := strings.Split(rule, ",")
	if len(parts) != 2 {
		panic("Invalid throttle format. Use throttle:60,1")
	}

	limit, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	window := time.Duration(minutes) * time.Minute

	return func(c *gin.Context) {
		limiter := getVisitor(c.ClientIP(), limit, window)
		allowed, remaining, retryAfter := limiter.allow()

		// Always send rate-limit headers
		c.Header("X-RateLimit-Limit", strconv.Itoa(limit))
		if allowed {
			c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
			c.Next()
		} else {
			c.Header("Retry-After", strconv.Itoa(retryAfter))
			c.Header("X-RateLimit-Remaining", "0")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}
	}
}

// Call this in main.go to start cleanup
func InitThrottleCleanup() {
	go cleanupVisitors()
}

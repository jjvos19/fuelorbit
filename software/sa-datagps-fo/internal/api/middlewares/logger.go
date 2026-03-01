package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()
		latency := time.Since(t)
		fmt.Printf("[Interceptor] Ruta: %s | Duracion: %v\n", ctx.Request.URL.Path, latency)
	}
}

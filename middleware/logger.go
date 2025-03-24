package middleware

import (
	"lsjv-nft-market-sync/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前的时间
		start := time.Now()

		c.Next()

		config.Logger.Info("访问日志",
			zap.String("请求方法", c.Request.Method),
			zap.String("请求路径", c.Request.URL.Path),
			zap.String("请求IP", c.ClientIP()),
			zap.Duration("请求耗时", time.Since(start)),
		)
	}
}

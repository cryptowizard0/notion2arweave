package service

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func respondJSONError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"requestID": requestid.Get(ctx),
		"code":      code,
		"message":   err.Error(),
	})
}

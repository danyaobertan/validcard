package delivery

import "github.com/gin-gonic/gin"

type (
	ValidCardHTTP interface {
		ValidateCardInfo() func(c *gin.Context)
	}
)

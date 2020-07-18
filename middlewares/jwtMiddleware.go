package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Print("BEFORE")

		c.Next()

		fmt.Print("DESPUES")
	}
}

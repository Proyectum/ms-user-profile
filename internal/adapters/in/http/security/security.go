package security

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

const space = " "
const bearer = "Bearer"

func JwtMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing authorization header",
		})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, space)
	if len(parts) != 2 || parts[0] != bearer {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bearer format is required",
		})
		c.Abort()
		return
	}

	jwtSecret := viper.GetString("security.jwt.secret")
	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown sign: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
				c.Abort()
				return
			}
		}

		c.Set("username", claims["username"])
		c.Set("email", claims["email"])
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
	c.Abort()
}

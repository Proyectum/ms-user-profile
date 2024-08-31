package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
)

const (
	headerAuthUser   = "X-Auth-User"
	headerAuthEmail  = "X-Auth-Email"
	headerAuthScopes = "X-Auth-Scopes"
	UsernameKey      = "username"
	EmailKey         = "email"
	ScopesKey        = "scopes"
	readScope        = "read"
	writeScope       = "write"
	adminScope       = "admin"
)

var readMethods = []string{http.MethodGet}
var writeMethods = []string{http.MethodPost, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch}

func AuthHeaderMiddleware(c *gin.Context) {
	user := c.GetHeader(headerAuthUser)
	email := c.GetHeader(headerAuthEmail)
	scopes := c.GetHeader(headerAuthScopes)

	if len(user) == 0 || len(email) == 0 || len(scopes) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing authorization headers",
		})
		c.Abort()
		return
	}

	scopesArr := strings.Split(scopes, ",")

	c.Set(UsernameKey, user)
	c.Set(EmailKey, email)
	c.Set(ScopesKey, scopesArr)
	c.Next()
}

func ScopesMiddleware(c *gin.Context) {
	method := c.Request.Method
	scopes := getStringSliceValue(c, ScopesKey)

	if slices.Contains(scopes, adminScope) {
		c.Next()
		return
	}

	if slices.Contains(readMethods, method) && slices.Contains(scopes, readScope) {
		c.Next()
		return
	}

	if slices.Contains(writeMethods, method) && slices.Contains(scopes, writeScope) {
		c.Next()
		return
	}

	c.JSON(http.StatusForbidden, gin.H{
		"message": "forbidden permissions",
	})
	c.Abort()
}

func IsAdmin(c *gin.Context) bool {
	scopes := getStringSliceValue(c, ScopesKey)
	return slices.Contains(scopes, adminScope)
}

func getStringSliceValue(c *gin.Context, key string) []string {
	if val, exists := c.Get(key); exists {
		return val.([]string)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
	c.Abort()
	return make([]string, 0)
}

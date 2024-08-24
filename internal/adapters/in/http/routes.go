package http

import (
	"github.com/gin-gonic/gin"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/api"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/security"
	"net/http"
)

type profileRoutes struct{}

func (p *profileRoutes) GetProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented yet",
	})
}

func (p *profileRoutes) UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented yet",
	})
}

func (p *profileRoutes) UpdateNotification(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented yet",
	})
}

func (p *profileRoutes) GetTypes(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented yet",
	})
}

func (p *profileRoutes) errorHandler(c *gin.Context, err error, status int) {

}

func RegisterRoutes(r *gin.Engine) {
	routes := profileRoutes{}
	api.RegisterHandlersWithOptions(
		r,
		&routes,
		api.GinServerOptions{
			Middlewares:  []api.MiddlewareFunc{security.JwtMiddleware},
			ErrorHandler: routes.errorHandler,
		})
}

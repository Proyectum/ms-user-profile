package http

import (
	"context"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/api"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/security"
	app "github.com/proyectum/ms-user-profile/internal/app/usecases"
	"github.com/proyectum/ms-user-profile/internal/domain/entities"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
	"net/http"
)

type profileRoutes struct {
	getUserProfileUseCase    usecases.GetUserProfileUseCase
	updateUserProfileUseCase usecases.UpdateUserProfileUseCase
}

func (p *profileRoutes) GetProfile(c *gin.Context) {
	username := getStringValue(c, "username")
	if c.IsAborted() {
		return
	}
	email := getStringValue(c, "email")
	if c.IsAborted() {
		return
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, "username", username)
	ctx = context.WithValue(ctx, "email", email)

	profile, err := p.getUserProfileUseCase.GetUserProfile(ctx)

	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, api.GeneralProfile{
		Bio:       &profile.Bio,
		Email:     openapi_types.Email(profile.Email),
		FirstName: profile.FirstName,
		Initials:  profile.Initials,
		LastName:  profile.LastName,
		Locale:    profile.Locale,
		Username:  profile.Username,
	})
}

func (p *profileRoutes) UpdateProfile(c *gin.Context) {
	username := getStringValue(c, "username")
	if c.IsAborted() {
		return
	}
	var update api.UpdateProfile

	if err := c.BindJSON(&update); err != nil {
		p.handleError(c, err)
		return
	}

	var entity = entities.UpdateProfile{
		LastName:  update.LastName,
		FirstName: update.FirstName,
		Locale:    update.Locale,
		Bio:       update.Bio,
	}
	err := p.updateUserProfileUseCase.Update(username, entity)

	if err != nil {
		p.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
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

func RegisterRoutes(r *gin.Engine) {
	routes := profileRoutes{
		getUserProfileUseCase:    app.NewUserProfileUseCase(),
		updateUserProfileUseCase: app.NewUpdateUserProfileUseCase(),
	}
	api.RegisterHandlersWithOptions(
		r,
		&routes,
		api.GinServerOptions{
			Middlewares: []api.MiddlewareFunc{security.JwtMiddleware},
		})
}

func getStringValue(c *gin.Context, key string) string {
	if val, exists := c.Get(key); exists {
		return val.(string)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
	c.Abort()
	return ""
}

func (p *profileRoutes) handleError(c *gin.Context, err error) {
	handleInternalError(err, c)
}

func handleInternalError(err error, c *gin.Context) {
	handleError(c, err, http.StatusInternalServerError)
}

func handleError(c *gin.Context, err error, status int) {
	c.JSON(status, gin.H{
		"code":    http.StatusText(status),
		"message": err.Error(),
	})
}

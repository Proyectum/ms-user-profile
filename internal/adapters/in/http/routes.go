package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/api"
	"github.com/proyectum/ms-user-profile/internal/adapters/in/http/security"
	app "github.com/proyectum/ms-user-profile/internal/app/usecases"
	"github.com/proyectum/ms-user-profile/internal/domain/usecases"
	"net/http"
)

type profileRoutes struct {
	getUserProfileUseCase      usecases.GetUserProfileUseCase
	updateUserProfileUseCase   usecases.UpdateUserProfileUseCase
	updateNotificationUseCase  usecases.UpdateNotificationUseCase
	getNotificationUseCase     usecases.GetNotificationUseCase
	getNotificationTypeUseCase usecases.GetNotificationTypeUseCase
	apiMapper                  UserProfileApiMapper
}

func (r *profileRoutes) GetNotifications(c *gin.Context, username string) {
	user := getStringValue(c, security.UsernameKey)
	if c.IsAborted() {
		return
	}

	r.checkPermissions(username, user, c)
	if c.IsAborted() {
		return
	}

	notifications, err := r.getNotificationUseCase.GetNotificationSettings(username)

	if err != nil {
		r.handleError(c, err)
		return
	}

	dtos := r.apiMapper.ToNotificationSettingDTOs(notifications)

	c.JSON(http.StatusOK, dtos)
}

func (r *profileRoutes) GetProfile(c *gin.Context, username string) {
	user := getStringValue(c, security.UsernameKey)
	if c.IsAborted() {
		return
	}

	r.checkPermissions(username, user, c)
	if c.IsAborted() {
		return
	}

	email := getStringValue(c, security.EmailKey)
	if c.IsAborted() {
		return
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, "username", username)
	ctx = context.WithValue(ctx, "email", email)

	profile, err := r.getUserProfileUseCase.GetUserProfile(ctx)

	if err != nil {
		r.handleError(c, err)
		return
	}

	dto := r.apiMapper.ToGeneralProfile(profile)

	c.JSON(http.StatusOK, dto)
}

func (r *profileRoutes) UpdateProfile(c *gin.Context, username string) {
	user := getStringValue(c, security.UsernameKey)
	if c.IsAborted() {
		return
	}

	r.checkPermissions(username, user, c)
	if c.IsAborted() {
		return
	}
	var update api.UpdateProfile

	if err := c.BindJSON(&update); err != nil {
		r.handleError(c, err)
		return
	}

	var entity = r.apiMapper.ToUpdateProfileDomain(&update)
	err := r.updateUserProfileUseCase.Update(username, *entity)

	if err != nil {
		r.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *profileRoutes) UpdateNotification(c *gin.Context, username string) {
	user := getStringValue(c, security.UsernameKey)
	if c.IsAborted() {
		return
	}

	r.checkPermissions(username, user, c)
	if c.IsAborted() {
		return
	}

	var update api.UpdateNotification

	if err := c.BindJSON(&update); err != nil {
		r.handleError(c, err)
		return
	}

	err := r.updateNotificationUseCase.Update(username, update.TypeId, update.Active)

	if err != nil {
		r.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *profileRoutes) GetTypes(c *gin.Context) {
	entities, err := r.getNotificationTypeUseCase.GetNotificationTypes()

	if err != nil {
		r.handleError(c, err)
		return
	}

	dtos := r.apiMapper.ToNotificationTypeDTOs(entities)

	c.JSON(http.StatusOK, dtos)
}

func RegisterRoutes(r *gin.Engine) {
	routes := profileRoutes{
		getUserProfileUseCase:      app.NewUserProfileUseCase(),
		updateUserProfileUseCase:   app.NewUpdateUserProfileUseCase(),
		updateNotificationUseCase:  app.NewUpdateNotificationUseCase(),
		getNotificationUseCase:     app.NewGetNotificationUseCase(),
		getNotificationTypeUseCase: app.NewGetNotificationTypeUseCase(),
		apiMapper:                  NewUserProfileApiMapper(),
	}
	api.RegisterHandlersWithOptions(
		r,
		&routes,
		api.GinServerOptions{
			Middlewares: []api.MiddlewareFunc{
				security.AuthHeaderMiddleware,
				security.ScopesMiddleware,
			},
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

func (r *profileRoutes) handleError(c *gin.Context, err error) {
	handleInternalError(err, c)
}

func (r *profileRoutes) checkPermissions(username string, user string, c *gin.Context) {
	if username == user {
		return
	}

	if security.IsAdmin(c) {
		return
	}

	c.Status(http.StatusForbidden)
	c.Abort()
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

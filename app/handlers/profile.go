package handlers

import (
	"context"
	"net/http"

	"github.com/frisk038/livechat/business/models"
	"github.com/gin-gonic/gin"
)

type business interface {
	CreateUser(ctx context.Context, user models.User) error
}

type HandlerProfile struct {
	business business
}

func NewHandlerProfile(b business) HandlerProfile {
	return HandlerProfile{business: b}
}

func (hp *HandlerProfile) PostUsers(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//TODO better fin tuning err
	if err := hp.business.CreateUser(ctx, user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
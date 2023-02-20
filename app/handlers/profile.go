package handlers

import (
	"context"
	"net/http"

	"github.com/frisk038/livechat/business/models"
	"github.com/gin-gonic/gin"
)

type business interface {
	CreateUser(ctx context.Context, user models.User) error
	SetHobbies(ctx context.Context, user models.User) error
}

type HandlerProfile struct {
	business business
}

func NewHandlerProfile(b business) HandlerProfile {
	return HandlerProfile{business: b}
}

func (hp *HandlerProfile) PostUsers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	ctx := c.Request.Context()
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//TODO better fine tuning err
	if err := hp.business.CreateUser(ctx, user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (hp *HandlerProfile) PostUsersHobbies(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	ctx := c.Request.Context()
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//TODO better fine tuning err
	if err := hp.business.SetHobbies(ctx, user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

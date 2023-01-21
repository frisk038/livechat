package main

import (
	"github.com/frisk038/livechat/app/handlers"
	"github.com/frisk038/livechat/business"
	"github.com/frisk038/livechat/infra/repo"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func initRoutes(hp handlers.HandlerProfile) {
	r := gin.Default()
	r.POST("/user", hp.PostUsers)

	r.Run(":8090")
}

func main() {
	repo, err := repo.NewRepo()
	if err != nil {
		log.Error(err)
	}

	bp := business.NewBusinessProfile(repo)
	hp := handlers.NewHandlerProfile(&bp)

	initRoutes(hp)
}

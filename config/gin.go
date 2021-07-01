package config

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Gin struct {
	GinEngine    *gin.Engine
	ServerConfig GinServer
}

func (gin Gin) StartServer() {
	log.Infof("Start Gin Server -%s:%d", gin.ServerConfig.hostIPAddress, gin.ServerConfig.hostPort)
}

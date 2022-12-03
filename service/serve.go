package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartServe() {
	router := gin.Default()
	router.Use(gin.Recovery())

	// path
	group := router.Group("/v1/")
	group.GET("/page/save/:uuid", SavePage2Ar)
	group.GET("/page/load/:parent/:artxid", LoadPageFromAr)

	port := fmt.Sprintf(":%s", viper.GetString("service.port"))
	router.Run(port)
}

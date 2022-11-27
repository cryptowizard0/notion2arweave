package service

import "github.com/gin-gonic/gin"

func StartServe() {
	router := gin.Default()
	router.Use(gin.Recovery())

	// path
	group := router.Group("/v1/")
	group.GET("/page/save/:uuid", SavePage2Ar)
	group.GET("/page/load/:parent/:artxid", LoadPageFromAr)

	router.Run(":2333")
}

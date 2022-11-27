package service

import (
	"errors"
	"net/http"

	"github.com/cryptowizard0/notion2arweave/operator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SavePage2Ar(c *gin.Context) {
	uuid := c.Param("uuid")
	log.WithField("uuid", uuid).Debug("Get request SavePage2Ar")

	opt := operator.CreateOperator(
		viper.GetString("notion.api_auth"),
		viper.GetString("arweave.pk"),
		"USDC")
	if opt == nil {
		log.Error("Error, create operator failed!")
		respondJSONError(c, http.StatusBadRequest, errors.New("server error"))
		return
	}

	arTxId, err := opt.SavePage2Ar(uuid)
	if err != nil {
		log.WithContext(WithGinContext(c)).Error("save page failed: ", err.Error())
		respondJSONError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"artxid": arTxId,
		},
	})
}

func LoadPageFromAr(c *gin.Context) {
	parentId := c.Param("parent")
	artxid := c.Param("artxid")
	log.WithField("artxid", artxid).Debug("Get request LoadPageFromAr")

	opt := operator.CreateOperator(
		viper.GetString("notion.api_auth"),
		viper.GetString("arweave.pk"),
		"USDC")
	if opt == nil {
		log.Error("Error, create operator failed!")
		respondJSONError(c, http.StatusBadRequest, errors.New("server error"))
		return
	}

	uuid, err := opt.LoadPageFromAr(parentId, artxid)
	if err != nil {
		log.WithContext(WithGinContext(c)).Error("load page  failed: ", err.Error())
		respondJSONError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"uuid": uuid,
		},
	})
}

package utils

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func LogResp_Error(resp *resty.Response) {
	log.Errorf("Response Status Code: %v", resp.StatusCode())
	log.Errorf("Response Status: %v", resp.Status())
	log.Errorf("Response Body: %v", resp)
	log.Errorf("Response Time: %v", resp.Time())
	log.Errorf("Response Received At: %v \n", resp.ReceivedAt())
}

func LogResp_Debug(resp *resty.Response) {
	log.Debugf("Response Status Code: %v", resp.StatusCode())
	log.Debugf("Response Status: %v", resp.Status())
	log.Debugf("Response Body: %v", resp)
	log.Debugf("Response Time: %v", resp.Time())
	log.Debugf("Response Received At: %v \n", resp.ReceivedAt())
}

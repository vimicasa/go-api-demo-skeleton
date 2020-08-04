package app

import (
	"go-api-demo-skeleton/config"

	"github.com/sirupsen/logrus"
)

var (
	// AppConf is gorush config
	AppConf config.ConfYaml
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
)

// Response generic response of the app
type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

package app

import (
	"github.com/vimicasa/go-api-demo-skeleton/config"

	"github.com/sirupsen/logrus"
)

var (
	// AppConf is gorush config
	AppConf config.ConfYaml
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogServer is log server error log
	LogServer *logrus.Logger
)

// Response generic response of the app
type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

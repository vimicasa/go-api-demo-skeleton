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

package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

const (
	// TODO: change this value
	// NameApp of the application
	NameApp = "test"
)

var defaultConf = []byte(`
core:
  address: "" # ip address to bind (default: any)
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"
  cert_base64: ""
  key_base64: ""
  auto_tls:
    enabled: false # Automatically install TLS certificates from Let's Encrypt.
    folder: ".cache" # folder for storing TLS certificates
    host: "" # which domains the Let's Encrypt will attempt

api:
  health_uri: "/healthz"

log:
  format: "string" # string or json
  access_log: "access.log" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "error.log" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"

`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core SectionCore `yaml:"core"`
	API  SectionAPI  `yaml:"api"`
	Log  SectionLog  `yaml:"log"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Address    string         `yaml:"address"`
	Port       string         `yaml:"port"`
	Mode       string         `yaml:"mode"`
	SSL        bool           `yaml:"ssl"`
	CertPath   string         `yaml:"cert_path"`
	KeyPath    string         `yaml:"key_path"`
	CertBase64 string         `yaml:"cert_base64"`
	KeyBase64  string         `yaml:"key_base64"`
	AutoTLS    SectionAutoTLS `yaml:"auto_tls"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
}

// SectionAutoTLS support Let's Encrypt setting.
type SectionAutoTLS struct {
	Enabled bool   `yaml:"enabled"`
	Folder  string `yaml:"folder"`
	Host    string `yaml:"host"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	HealthURI string `yaml:"health_uri"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()        // read in environment variables that match
	viper.SetEnvPrefix(NameApp) // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory
		viper.AddConfigPath("/etc/app/")
		viper.AddConfigPath("$HOME/." + NameApp)
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Address = viper.GetString("core.address")
	conf.Core.Port = viper.GetString("core.port")
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.SSL = viper.GetBool("core.ssl")
	conf.Core.CertPath = viper.GetString("core.cert_path")
	conf.Core.KeyPath = viper.GetString("core.key_path")
	conf.Core.CertBase64 = viper.GetString("core.cert_base64")
	conf.Core.KeyBase64 = viper.GetString("core.key_base64")
	conf.Core.AutoTLS.Enabled = viper.GetBool("core.auto_tls.enabled")
	conf.Core.AutoTLS.Folder = viper.GetString("core.auto_tls.folder")
	conf.Core.AutoTLS.Host = viper.GetString("core.auto_tls.host")

	// Api
	conf.API.HealthURI = viper.GetString("api.health_uri")

	// log
	conf.Log.Format = viper.GetString("log.format")
	conf.Log.AccessLog = viper.GetString("log.access_log")
	conf.Log.AccessLevel = viper.GetString("log.access_level")
	conf.Log.ErrorLog = viper.GetString("log.error_log")
	conf.Log.ErrorLevel = viper.GetString("log.error_level")

	return conf, nil
}

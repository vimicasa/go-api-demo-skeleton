package router

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"go-api-demo-skeleton/app"
	h "go-api-demo-skeleton/app/handler"
	"go-api-demo-skeleton/app/router/middleware"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
)

var (
	rxURL = regexp.MustCompile(`^/healthz$`)
)

// RunHTTPServer loads the middlewares, routes, handlers.
func RunHTTPServer() (err error) {

	server := &http.Server{
		Addr:    app.AppConf.Core.Address + ":" + app.AppConf.Core.Port,
		Handler: routerEngine(),
	}

	app.LogAccess.Info("HTTPD server is running on " + app.AppConf.Core.Port + " port.")
	if app.AppConf.Core.AutoTLS.Enabled {
		return startServer(autoTLSServer())
	} else if app.AppConf.Core.SSL {
		config := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		if config.NextProtos == nil {
			config.NextProtos = []string{"http/1.1"}
		}

		config.Certificates = make([]tls.Certificate, 1)
		if app.AppConf.Core.CertPath != "" && app.AppConf.Core.KeyPath != "" {
			config.Certificates[0], err = tls.LoadX509KeyPair(app.AppConf.Core.CertPath, app.AppConf.Core.KeyPath)
			if err != nil {
				app.LogError.Error("Failed to load https cert file: ", err)
				return err
			}
		} else if app.AppConf.Core.CertBase64 != "" && app.AppConf.Core.KeyBase64 != "" {
			cert, err := base64.StdEncoding.DecodeString(app.AppConf.Core.CertBase64)
			if err != nil {
				app.LogError.Error("base64 decode error:", err.Error())
				return err
			}
			key, err := base64.StdEncoding.DecodeString(app.AppConf.Core.KeyBase64)
			if err != nil {
				app.LogError.Error("base64 decode error:", err.Error())
				return err
			}
			if config.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
				app.LogError.Error("tls key pair error:", err.Error())
				return err
			}
		} else {
			return errors.New("missing https cert config")
		}

		server.TLSConfig = config
	}

	return startServer(server)

}

func autoTLSServer() *http.Server {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(app.AppConf.Core.AutoTLS.Host),
		Cache:      autocert.DirCache(app.AppConf.Core.AutoTLS.Folder),
	}

	return &http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		Handler:   routerEngine(),
	}
}

func startServer(s *http.Server) error {
	if s.TLSConfig == nil {
		return s.ListenAndServe()
	}

	return s.ListenAndServeTLS("", "")
}

func routerEngine() *gin.Engine {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if app.AppConf.Core.Mode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	if app.IsTerm {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	}

	// set server mode
	gin.SetMode(app.AppConf.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(logger.SetLogger(logger.Config{
		UTC:            true,
		SkipPathRegexp: rxURL,
	}))
	r.Use(gin.Recovery())
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestID())

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	api := r.Group("/")
	// no authentication endpoints
	{
		api.GET(app.AppConf.API.HealthURI, h.HeartbeatHandler)
		api.HEAD(app.AppConf.API.HealthURI, h.HeartbeatHandler)
		api.GET("/version", h.VersionHandler)

		api.POST("/login", h.LoginHandler)
		api.GET("/", rootHandler)
	}
	// basic authentication endpoints
	readAuth := r.Group("/")
	{
		readAuth.Use(middleware.AuthenticationRequired())
		{
			readAuth.GET("/read", h.LogoutHandler)
		}
	}
	// admin authentication endpoints
	basicAuth := r.Group("/")
	{
		basicAuth.Use(middleware.AuthenticationRequired("admin", "basic"))
		{
			basicAuth.GET("/basic", rootHandler)
		}
	}
	// admin authentication endpoints
	adminAuth := r.Group("/")
	{
		adminAuth.Use(middleware.AuthenticationRequired("admin"))
		{
			adminAuth.GET("/admin", rootHandler)
		}
	}

	return r
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome Root Handler:" + c.Request.URL.Path,
	})
}

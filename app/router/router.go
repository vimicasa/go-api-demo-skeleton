package router

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net/http"
	"regexp"

	"github.com/vimicasa/go-api-demo-skeleton/app"
	h "github.com/vimicasa/go-api-demo-skeleton/app/handler"
	"github.com/vimicasa/go-api-demo-skeleton/app/router/middleware"

	"github.com/gin-gonic/gin"
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

	app.LogServer.Info("HTTPD server is running on " + app.AppConf.Core.Port + " port.")
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
				app.LogServer.Error("Failed to load https cert file: ", err)
				return err
			}
		} else if app.AppConf.Core.CertBase64 != "" && app.AppConf.Core.KeyBase64 != "" {
			cert, err := base64.StdEncoding.DecodeString(app.AppConf.Core.CertBase64)
			if err != nil {
				app.LogServer.Error("base64 decode error:", err.Error())
				return err
			}
			key, err := base64.StdEncoding.DecodeString(app.AppConf.Core.KeyBase64)
			if err != nil {
				app.LogServer.Error("base64 decode error:", err.Error())
				return err
			}
			if config.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
				app.LogServer.Error("tls key pair error:", err.Error())
				return err
			}
		} else {
			return errors.New("missing https cert config")
		}

		server.TLSConfig = config
	}

	return startServer(server)
}

func routerEngine() *gin.Engine {

	// set server mode
	gin.SetMode(app.AppConf.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(middleware.Logger())
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

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome Root Handler:" + c.Request.URL.Path,
	})
}
